/*
 * Copyright 2007-2021 Charles du Jeu - Abstrium SAS <team (at) pyd.io>
 * This file is part of Pydio.
 *
 * Pydio is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */
import React from 'react'
import Pydio from 'pydio'
import emptyDataModel from "./emptyDataModel";
import {debounce} from 'lodash';
import SearchApi from 'pydio/http/search-api'


export default function withSearch(Component, historyIdentifier, scope){

    return class WithSearch extends React.Component {

        constructor(props) {
            super(props);
            this.performSearch = debounce(this.performSearch.bind(this), 500)
            let {values = {}} = props;
            values = {scope: scope || props.scope || 'folder', ...values};
            this.state = {
                dataModel: props.dataModel || props.pydio.getContextHolder() || emptyDataModel(),
                values,
                limit: 30,
                empty: true,
                loading: false,
                facets:[],
                activeFacets:[],
                history: []
            }
            if(historyIdentifier){
                this.loadHistory().then(hh => this.setState({history: hh}))
            }
        }

        componentDidUpdate(prevProps, prevState, snapshot) {
            if(prevState !== this.state) {
                const {onSearchStateChange} = this.props;
                if(onSearchStateChange) {
                    onSearchStateChange(this.state);
                }
            }
        }

        performSearch() {
            const {values, limit, dataModel, activeFacets} = this.state;
            const searchRootNode = dataModel.getSearchNode();
            searchRootNode.getMetadata().set('search_values', values);
            searchRootNode.getMetadata().set('active_facets', activeFacets);
            searchRootNode.observeOnce("loaded", ()=> {
                dataModel.setContextNode(searchRootNode, true);
            })
            searchRootNode.setChildren([]);
            searchRootNode.setLoaded(false);
            const {scope, ...searchValues} = values;

            const keys = Object.keys(searchValues);
            if (keys.length === 0 || (keys.length === 1 && keys[0] === 'basenameOrContent' && !values['basenameOrContent'])) {
                searchRootNode.clear();
                searchRootNode.setLoaded(true);
                searchRootNode.notify("loaded");
                this.setState({loading: false,empty: true, facets:[], activeFacets:[]});
                return;
            }

            this.setState({loading: true, empty: false});
            searchRootNode.setLoading(true);
            searchRootNode.notify("loading");
            const api = new SearchApi(Pydio.getInstance());
            api.search(this.mergeFacets(values, activeFacets), scope, limit).then(response => {
                searchRootNode.clear();
                const res = response.Results || []
                res.forEach(node => {
                    searchRootNode.addChild(node)
                })
                searchRootNode.setLoading(false);
                searchRootNode.setLoaded(true);
                searchRootNode.notify("loaded");
                if(historyIdentifier && values.basenameOrContent) {
                    this.pushHistory(values.basenameOrContent)
                }
                this.setState({
                    loading: false,
                    facets: response.Facets ||[]
                });
            }).catch(()=>{
                searchRootNode.clear();
                searchRootNode.setLoading(false);
                searchRootNode.notify("loaded");
                this.setState({loading: false});
            });

        }
        
        setValues(newValues){
            const {onUpdateSearch} = this.props;
            Object.keys(newValues).forEach(k => {
                if(!newValues[k]){
                    delete(newValues[k])
                }
            });
            let {scope, ...other} = newValues;
            if(!scope) {
                const {values} = this.state;
                scope = values.scope;
            }
            this.setState({values: {scope, ...other}, facets:[], activeFacets:[]}, this.performSearch);
            if(onUpdateSearch){
                onUpdateSearch({values: newValues});
            }
        }

        setLimit(limit){
            this.setState({limit}, this.performSearch);
        }

        toggleFacet(facet, toggle){
            const {activeFacets = []} = this.state;
            let newFacets = []
            if(toggle){
                newFacets = [...activeFacets, facet];
            } else {
                newFacets = activeFacets.filter(s => !(s.FieldName===facet.FieldName && s.Label === facet.Label))
            }
            this.setState({activeFacets:newFacets}, this.performSearch)
        }

        mergeFacets(values = {}, facets = []){
            let data = {};
            const {basenameOrContent} = values;
            facets.forEach(facet => {
                switch (facet.FieldName){
                    case "Size":
                        data['ajxp_bytesize'] = {from:facet.Min, to:facet.Max}
                        break;
                    case "ModifTime":
                        data['ajxp_modiftime'] = {from:facet.Start*1000, to:facet.End*1000}
                        break;
                    case "Extension":
                        data['ajxp_mime'] = facet.Label
                        break;
                    case "NodeType":
                        data['ajxp_mime'] = 'ajxp_' + facet.Label
                        break;
                    case "TextContent":
                        data['basenameOrContent'] = ''
                        data['Content'] = basenameOrContent
                        break;
                    case "Basename":
                        data['basenameOrContent'] = ''
                        data['basename'] = basenameOrContent
                        break;
                    default:
                        if(facet.FieldName.indexOf('Meta.') === 0) {
                            data['ajxp_meta_' + facet.FieldName.replace('Meta.', '')] = facet.Label;
                        }
                        break;
                }
            })
            return {...values, ...data};
        }


        getUserHistoryKey(){
            return Pydio.getInstance().user.getIdmUser().then(
                u => "cells.search-engine." + historyIdentifier + "." + u.Uuid
            );
        }

        loadHistory(){
            return this.getUserHistoryKey().then(key => {
                const i = localStorage.getItem(key)
                if(!i) {
                    return []
                }
                try {
                    const data = JSON.parse(i)
                    if(data.map){
                        return data;
                    }
                    return [];
                }catch (e){
                    return []
                }
            })
        }

        pushHistory(term){
            if(!term){
                return
            }
            const {history = []} = this.state;
            const newHistory = history.filter(f => f !== term).slice(0, 19); // store only 20 terms
            newHistory.unshift(term);
            this.getUserHistoryKey().then(key => {
                this.setState({history: newHistory}, () => {
                    localStorage.setItem(key, JSON.stringify(newHistory));
                })
            })
        }

        render() {
            return (
                <Component 
                    {...this.props} 
                    {...this.state}
                    submitSearch={this.performSearch}
                    setValues={this.setValues.bind(this)}
                    setLimit={this.setLimit.bind(this)}
                    toggleFacet={this.toggleFacet.bind(this)}
                />
            );
        }

    }

}