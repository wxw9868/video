!function(e,t,n,o,r){var i="undefined"!=typeof globalThis?globalThis:"undefined"!=typeof self?self:"undefined"!=typeof window?window:"undefined"!=typeof global?global:{},l="function"==typeof i[o]&&i[o],u=l.cache||{},f="undefined"!=typeof module&&"function"==typeof module.require&&module.require.bind(module);function d(t,n){if(!u[t]){if(!e[t]){var r="function"==typeof i[o]&&i[o];if(!n&&r)return r(t,!0);if(l)return l(t,!0);if(f&&"string"==typeof t)return f(t);var a=new Error("Cannot find module '"+t+"'");throw a.code="MODULE_NOT_FOUND",a}s.resolve=function(n){var o=e[t][1][n];return null!=o?o:n},s.cache={};var c=u[t]=new d.Module(t);e[t][0].call(c.exports,s,c,c.exports,this)}return u[t].exports;function s(e){var t=s.resolve(e);return!1===t?{}:d(t)}}d.isParcelRequire=!0,d.Module=function(e){this.id=e,this.bundle=d,this.exports={}},d.modules=e,d.cache=u,d.parent=l,d.register=function(t,n){e[t]=[function(e,t){t.exports=n},{}]},Object.defineProperty(d,"root",{get:function(){return i[o]}}),i[o]=d;for(var a=0;a<t.length;a++)d(t[a]);if(n){var c=d(n);"object"==typeof exports&&"undefined"!=typeof module?module.exports=c:"function"==typeof define&&define.amd&&define((function(){return c}))}}({oLNvk:[function(e,t,n){e("@parcel/transformer-js/src/esmodule-helpers.js").defineInteropFlag(n);const o={"Video Info":"統計訊息",Close:"關閉","Video Load Failed":"載入失敗",Volume:"音量",Play:"播放",Pause:"暫停",Rate:"速度",Mute:"靜音","Video Flip":"畫面翻轉",Horizontal:"水平",Vertical:"垂直",Reconnect:"重新連接","Show Setting":"顯示设置","Hide Setting":"隱藏设置",Screenshot:"截圖","Play Speed":"播放速度","Aspect Ratio":"畫面比例",Default:"默認",Normal:"正常",Open:"打開","Switch Video":"切換","Switch Subtitle":"切換字幕",Fullscreen:"全屏","Exit Fullscreen":"退出全屏","Web Fullscreen":"網頁全屏","Exit Web Fullscreen":"退出網頁全屏","Mini Player":"迷你播放器","PIP Mode":"開啟畫中畫","Exit PIP Mode":"退出畫中畫","PIP Not Supported":"不支持畫中畫","Fullscreen Not Supported":"不支持全屏","Subtitle Offset":"字幕偏移","Last Seen":"上次看到","Jump Play":"跳轉播放",AirPlay:"隔空播放","AirPlay Not Available":"隔空播放不可用"};n.default=o,"undefined"!=typeof window&&(window["artplayer-i18n-zh-tw"]=o)},{"@parcel/transformer-js/src/esmodule-helpers.js":"gKGrl"}],gKGrl:[function(e,t,n){n.interopDefault=function(e){return e&&e.__esModule?e:{default:e}},n.defineInteropFlag=function(e){Object.defineProperty(e,"__esModule",{value:!0})},n.exportAll=function(e,t){return Object.keys(e).forEach((function(n){"default"===n||"__esModule"===n||t.hasOwnProperty(n)||Object.defineProperty(t,n,{enumerable:!0,get:function(){return e[n]}})})),t},n.export=function(e,t,n){Object.defineProperty(e,t,{enumerable:!0,get:n})}},{}]},["oLNvk"],"oLNvk","parcelRequire94c2");