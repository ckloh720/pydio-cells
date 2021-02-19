'use strict';

Object.defineProperty(exports, '__esModule', {
    value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _createReactClass = require('create-react-class');

var _createReactClass2 = _interopRequireDefault(_createReactClass);

var _cellsSdk = require('cells-sdk');

var _materialUi = require('material-ui');

var _materialUiStyles = require('material-ui/styles');

/*
 * Copyright 2007-2017 Charles du Jeu - Abstrium SAS <team (at) pyd.io>
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

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _pydio = require('pydio');

var _pydio2 = _interopRequireDefault(_pydio);

var _ServiceExposedConfigs = require('./ServiceExposedConfigs');

var _ServiceExposedConfigs2 = _interopRequireDefault(_ServiceExposedConfigs);

var _MailerTest = require('./MailerTest');

var _MailerTest2 = _interopRequireDefault(_MailerTest);

/**
 * Editor for a given plugin. By default, displays documentation in a left column panel,
 * and plugin parameters as form cards on the right.
 * May take additionalPanes to be appended to the form cards.
 */

var PydioForm = _pydio2['default'].requireLib("form");
var PluginEditor = (0, _createReactClass2['default'])({
    displayName: 'PluginEditor',
    mixins: [AdminComponents.MessagesConsumerMixin],

    propTypes: {
        serviceName: _propTypes2['default'].string,
        rootNode: _propTypes2['default'].instanceOf(AjxpNode).isRequired,
        close: _propTypes2['default'].func,
        style: _propTypes2['default'].string,
        className: _propTypes2['default'].string,
        additionalPanes: _propTypes2['default'].shape({
            top: _propTypes2['default'].array,
            bottom: _propTypes2['default'].array
        }),
        docAsAdditionalPane: _propTypes2['default'].bool,
        additionalDescription: _propTypes2['default'].string,
        registerCloseCallback: _propTypes2['default'].func,
        onBeforeSave: _propTypes2['default'].func,
        onAfterSave: _propTypes2['default'].func,
        onRevert: _propTypes2['default'].func,
        onDirtyChange: _propTypes2['default'].func
    },

    getInitialState: function getInitialState() {

        return {
            loaded: false,
            documentation: '',
            dirty: false,
            label: '',
            docOpen: false
        };
    },

    externalSetDirty: function externalSetDirty() {
        this.setState({ dirty: true });
    },

    save: function save() {
        this.refs.formConfigs.save();
        this.setState({ dirty: false });
    },

    revert: function revert() {
        this.refs.formConfigs.revert();
        this.setState({ dirty: false });
    },

    parameterHasHelper: function parameterHasHelper(paramName, testPluginId) {
        var parameterName = paramName.split('/').pop();
        var h = PydioForm.Manager.hasHelper(this.props.serviceName, parameterName);
        if (!h && testPluginId) {
            h = PydioForm.Manager.hasHelper(testPluginId, parameterName);
        }
        return h;
    },

    showHelper: function showHelper(helperData, testPluginId) {
        if (helperData) {
            var serviceName = this.props.serviceName;
            if (testPluginId && !PydioForm.Manager.hasHelper(serviceName, helperData['name'])) {
                helperData['pluginId'] = testPluginId;
            } else {
                helperData['pluginId'] = serviceName;
            }
            helperData['updateCallback'] = this.helperUpdateValues.bind(this);
        }
        this.setState({ helperData: helperData });
    },

    closeHelper: function closeHelper() {
        this.setState({ helperData: null });
    },

    /**
     * External helper can pass a full set of values and update them
     * @param newValues
     */
    helperUpdateValues: function helperUpdateValues(newValues) {
        this.onChange(newValues, true);
    },

    toggleDocPane: function toggleDocPane() {
        this.setState({ docOpen: !this.state.docOpen });
    },

    monitorMainPaneScrolling: function monitorMainPaneScrolling(event) {
        if (event.target.className.indexOf('pydio-form-panel') === -1) {
            return;
        }
        var scroll = event.target.scrollTop;
        var newState = scroll > 5;
        var currentScrolledState = this.state && this.state.mainPaneScrolled;
        if (newState !== currentScrolledState) {
            this.setState({ mainPaneScrolled: newState });
        }
    },

    render: function render() {
        var _this = this;

        var _props = this.props;
        var additionalPanes = _props.additionalPanes;
        var closeEditor = _props.closeEditor;
        var docAsAdditionalPane = _props.docAsAdditionalPane;
        var className = _props.className;
        var style = _props.style;
        var rootNode = _props.rootNode;
        var tabs = _props.tabs;
        var accessByName = _props.accessByName;
        var muiTheme = _props.muiTheme;
        var _state = this.state;
        var documentation = _state.documentation;
        var pluginId = _state.pluginId;
        var docOpen = _state.docOpen;
        var mainPaneScrolled = _state.mainPaneScrolled;
        var dirty = _state.dirty;
        var parameters = _state.parameters;
        var values = _state.values;
        var helperData = _state.helperData;

        var adminStyles = AdminComponents.AdminStyles(muiTheme.palette);

        var addPanes = { top: [], bottom: [] };
        if (additionalPanes) {
            addPanes.top = additionalPanes.top.slice();
            addPanes.bottom = additionalPanes.bottom.slice();
        }
        var serviceName = this.props.serviceName;

        if (serviceName === 'pydio.grpc.mailer') {
            addPanes.bottom.push(_react2['default'].createElement(_MailerTest2['default'], { pydio: this.props.pydio, adminStyles: adminStyles }));
        }

        var closeButton = undefined;
        if (closeEditor) {
            closeButton = _react2['default'].createElement(_materialUi.RaisedButton, { label: this.context.getMessage('86', ''), onTouchTap: closeEditor });
        }

        var doc = documentation;
        if (doc && docAsAdditionalPane) {
            doc = doc.firstChild.nodeValue.replace('<p><ul', '<ul').replace('</ul></p>', '</ul>').replace('<p></p>', '');
            doc = doc.replace('<img src="', '<img style="width:90%;" src="plug/' + pluginId + '/');
            var readDoc = function readDoc() {
                return { __html: doc };
            };
            addPanes.top.push(_react2['default'].createElement(
                'div',
                { className: "plugin-doc" + (docOpen ? ' plugin-doc-open' : '') },
                _react2['default'].createElement(
                    'h3',
                    null,
                    'Documentation'
                ),
                _react2['default'].createElement('div', { className: 'plugin-doc-pane', dangerouslySetInnerHTML: readDoc() })
            ));
        }

        var scrollingClassName = '';
        if (this.state && mainPaneScrolled) {
            scrollingClassName = ' main-pane-scrolled';
        }
        var actions = [];
        if (accessByName('Create')) {
            var props = adminStyles.props.header.flatButton;
            if (!dirty) {
                props = adminStyles.props.header.flatButtonDisabled;
            }
            actions.push(_react2['default'].createElement(_materialUi.FlatButton, _extends({ secondary: true, disabled: !dirty, label: this.context.getMessage('plugins.6'), onTouchTap: this.revert }, props)));
            actions.push(_react2['default'].createElement(_materialUi.FlatButton, _extends({ secondary: true, disabled: !dirty, label: this.context.getMessage('plugins.5'), onTouchTap: this.save }, props)));
        }
        actions.push(closeButton);

        var icon = rootNode.getMetadata().get('icon_class');
        var label = rootNode.getLabel();
        // Building  a form
        return _react2['default'].createElement(
            'div',
            { className: (className ? className + " " : "") + "main-layout-nav-to-stack vertical-layout plugin-board" + scrollingClassName, style: style },
            _react2['default'].createElement(AdminComponents.Header, { title: label, actions: actions, scrolling: this.state && mainPaneScrolled, icon: icon }),
            _react2['default'].createElement(_ServiceExposedConfigs2['default'], _extends({
                ref: "formConfigs"
            }, this.props, {
                additionalPanes: addPanes,
                tabs: tabs,
                setHelperData: this.showHelper,
                checkHasHelper: this.parameterHasHelper,
                onScrollCallback: this.monitorMainPaneScrolling,
                className: 'row-flex',
                onDirtyChange: function (dirty) {
                    _this.setState({ dirty: dirty });
                }
            })),
            _react2['default'].createElement(PydioForm.PydioHelper, {
                helperData: helperData,
                close: this.closeHelper
            }),
            adminStyles.formCss()
        );
    }
});

exports['default'] = PluginEditor = (0, _materialUiStyles.muiThemeable)()(PluginEditor);

exports['default'] = PluginEditor;
module.exports = exports['default'];
