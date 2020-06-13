/*!
 * GraphQL Voyager - Represent any GraphQL API as an interactive graph
 * -------------------------------------------------------------
 *   Version: "1.0.0-rc.30"
 *   Repo: https://github.com/APIs-guru/graphql-voyager
 */
(function webpackUniversalModuleDefinition(root, factory) {
	if(typeof exports === 'object' && typeof module === 'object')
		module.exports = factory();
	else if(typeof define === 'function' && define.amd)
		define("GraphQLVoyager", [], factory);
	else if(typeof exports === 'object')
		exports["GraphQLVoyager"] = factory();
	else
		root["GraphQLVoyager"] = factory();
})(window, function() {
return /******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, { enumerable: true, get: getter });
/******/ 		}
/******/ 	};
/******/
/******/ 	// define __esModule on exports
/******/ 	__webpack_require__.r = function(exports) {
/******/ 		if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 			Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 		}
/******/ 		Object.defineProperty(exports, '__esModule', { value: true });
/******/ 	};
/******/
/******/ 	// create a fake namespace object
/******/ 	// mode & 1: value is a module id, require it
/******/ 	// mode & 2: merge all properties of value into the ns
/******/ 	// mode & 4: return value when already ns object
/******/ 	// mode & 8|1: behave like require
/******/ 	__webpack_require__.t = function(value, mode) {
/******/ 		if(mode & 1) value = __webpack_require__(value);
/******/ 		if(mode & 8) return value;
/******/ 		if((mode & 4) && typeof value === 'object' && value && value.__esModule) return value;
/******/ 		var ns = Object.create(null);
/******/ 		__webpack_require__.r(ns);
/******/ 		Object.defineProperty(ns, 'default', { enumerable: true, value: value });
/******/ 		if(mode & 2 && typeof value != 'string') for(var key in value) __webpack_require__.d(ns, key, function(key) { return value[key]; }.bind(null, key));
/******/ 		return ns;
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 24);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ (function(module, exports) {

module.exports = require("react");

/***/ }),
/* 1 */
/***/ (function(module, exports) {

module.exports = require("lodash");

/***/ }),
/* 2 */
/***/ (function(module, exports) {

module.exports = require("classnames");

/***/ }),
/* 3 */
/***/ (function(module, exports) {

module.exports = require("prop-types");

/***/ }),
/* 4 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/IconButton");

/***/ }),
/* 5 */
/***/ (function(module, exports) {

module.exports = {
	"monospaceFontFamily": "'Consolas', 'Inconsolata', 'Droid Sans Mono', 'Monaco', monospace",
	"baseFontFamily": "'helvetica neue', helvetica, arial, sans-serif",
	"baseFontSize": 14,
	"spacingUnit": 5,
	"panelItemsSpacing": 8,
	"panelSpacing": 15,
	"iconsSize": 24,
	"primaryColor": "#00bcd4",
	"backgroundColor": "#fff",
	"darkBgColor": "#0b2840",
	"highlightColor": "#00bcd4",
	"secondaryColor": "#548f9e",
	"logoColor": "#27535e",
	"linkColor": "#42a0dd",
	"linkHoverColor": "#0262a0",
	"fieldNameColor": "#224d6f",
	"builtinColor": "#711c1c",
	"textColor": "#666",
	"shadowColor": "rgba(0, 0, 0, .1)",
	"alertColor": "#b71c1c",
	"modalBgColor": "#0b2840",
	"docPanelWidth": 320,
	"typeInfoPopoverWidth": 320,
	"docPanelBgColor": "#fff",
	"docPanelItemStripeColor": "rgba(158, 158, 158, .07)",
	"docPanelItemHoverColor": "rgba(214, 236, 238, .6)",
	"argDefaultColor": "#0B7FC7",
	"argNameColor": "#c77f53",
	"nodeFillColor": "#f6f8f8",
	"nodeHeaderColor": "#548f9e",
	"nodeHeaderTextColor": "white",
	"edgeColor": "rgb(56, 97, 107)",
	"selectedEdgeColor": "red",
	"selectedFieldBg": "rgba(255, 0, 0, .18)",
	"smallViewport": "(max-width: 780px)",
	"bigViewport": "(min-width: 781px)"
}

/***/ }),
/* 6 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/Checkbox");

/***/ }),
/* 7 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/styles");

/***/ }),
/* 8 */
/***/ (function(module, exports) {

module.exports = require("graphql");

/***/ }),
/* 9 */
/***/ (function(module, exports, __webpack_require__) {

/**
 * Modules
 */

var elapsed = __webpack_require__(26)
var tween = __webpack_require__(28)
var raf = __webpack_require__(30)

/**
 * Constants
 */

var fps60 = 1000 / 60

/**
 * Expose animate
 */

module.exports = animate

/**
 * animate
 */

function animate (start, end, render, duration, easing) {
  var tick = tween(start, end, duration, easing, fps60)
  var time = elapsed()

  var id = raf(function ticker () {
    var frame = tick(time() / fps60)
    render(frame)

    if (frame !== end) {
      id = raf(ticker)
    }
  })

  return function () {
    raf.cancel(id)
  }
}


/***/ }),
/* 10 */
/***/ (function(module, exports) {

module.exports = require("commonmark");

/***/ }),
/* 11 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/MenuItem");

/***/ }),
/* 12 */
/***/ (function(module, exports) {

module.exports = require("react-dom");

/***/ }),
/* 13 */
/***/ (function(module, exports) {

module.exports = require("graphql/utilities");

/***/ }),
/* 14 */
/***/ (function(module, exports) {

module.exports = require("svg-pan-zoom");

/***/ }),
/* 15 */
/***/ (function(module, exports, __webpack_require__) {

/* WEBPACK VAR INJECTION */(function(process) {// .dirname, .basename, and .extname methods are extracted from Node.js v8.11.1,
// backported and transplited with Babel, with backwards-compat fixes

// Copyright Joyent, Inc. and other Node contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.

// resolves . and .. elements in a path array with directory names there
// must be no slashes, empty elements, or device names (c:\) in the array
// (so also no leading and trailing slashes - it does not distinguish
// relative and absolute paths)
function normalizeArray(parts, allowAboveRoot) {
  // if the path tries to go above the root, `up` ends up > 0
  var up = 0;
  for (var i = parts.length - 1; i >= 0; i--) {
    var last = parts[i];
    if (last === '.') {
      parts.splice(i, 1);
    } else if (last === '..') {
      parts.splice(i, 1);
      up++;
    } else if (up) {
      parts.splice(i, 1);
      up--;
    }
  }

  // if the path is allowed to go above the root, restore leading ..s
  if (allowAboveRoot) {
    for (; up--; up) {
      parts.unshift('..');
    }
  }

  return parts;
}

// path.resolve([from ...], to)
// posix version
exports.resolve = function() {
  var resolvedPath = '',
      resolvedAbsolute = false;

  for (var i = arguments.length - 1; i >= -1 && !resolvedAbsolute; i--) {
    var path = (i >= 0) ? arguments[i] : process.cwd();

    // Skip empty and invalid entries
    if (typeof path !== 'string') {
      throw new TypeError('Arguments to path.resolve must be strings');
    } else if (!path) {
      continue;
    }

    resolvedPath = path + '/' + resolvedPath;
    resolvedAbsolute = path.charAt(0) === '/';
  }

  // At this point the path should be resolved to a full absolute path, but
  // handle relative paths to be safe (might happen when process.cwd() fails)

  // Normalize the path
  resolvedPath = normalizeArray(filter(resolvedPath.split('/'), function(p) {
    return !!p;
  }), !resolvedAbsolute).join('/');

  return ((resolvedAbsolute ? '/' : '') + resolvedPath) || '.';
};

// path.normalize(path)
// posix version
exports.normalize = function(path) {
  var isAbsolute = exports.isAbsolute(path),
      trailingSlash = substr(path, -1) === '/';

  // Normalize the path
  path = normalizeArray(filter(path.split('/'), function(p) {
    return !!p;
  }), !isAbsolute).join('/');

  if (!path && !isAbsolute) {
    path = '.';
  }
  if (path && trailingSlash) {
    path += '/';
  }

  return (isAbsolute ? '/' : '') + path;
};

// posix version
exports.isAbsolute = function(path) {
  return path.charAt(0) === '/';
};

// posix version
exports.join = function() {
  var paths = Array.prototype.slice.call(arguments, 0);
  return exports.normalize(filter(paths, function(p, index) {
    if (typeof p !== 'string') {
      throw new TypeError('Arguments to path.join must be strings');
    }
    return p;
  }).join('/'));
};


// path.relative(from, to)
// posix version
exports.relative = function(from, to) {
  from = exports.resolve(from).substr(1);
  to = exports.resolve(to).substr(1);

  function trim(arr) {
    var start = 0;
    for (; start < arr.length; start++) {
      if (arr[start] !== '') break;
    }

    var end = arr.length - 1;
    for (; end >= 0; end--) {
      if (arr[end] !== '') break;
    }

    if (start > end) return [];
    return arr.slice(start, end - start + 1);
  }

  var fromParts = trim(from.split('/'));
  var toParts = trim(to.split('/'));

  var length = Math.min(fromParts.length, toParts.length);
  var samePartsLength = length;
  for (var i = 0; i < length; i++) {
    if (fromParts[i] !== toParts[i]) {
      samePartsLength = i;
      break;
    }
  }

  var outputParts = [];
  for (var i = samePartsLength; i < fromParts.length; i++) {
    outputParts.push('..');
  }

  outputParts = outputParts.concat(toParts.slice(samePartsLength));

  return outputParts.join('/');
};

exports.sep = '/';
exports.delimiter = ':';

exports.dirname = function (path) {
  if (typeof path !== 'string') path = path + '';
  if (path.length === 0) return '.';
  var code = path.charCodeAt(0);
  var hasRoot = code === 47 /*/*/;
  var end = -1;
  var matchedSlash = true;
  for (var i = path.length - 1; i >= 1; --i) {
    code = path.charCodeAt(i);
    if (code === 47 /*/*/) {
        if (!matchedSlash) {
          end = i;
          break;
        }
      } else {
      // We saw the first non-path separator
      matchedSlash = false;
    }
  }

  if (end === -1) return hasRoot ? '/' : '.';
  if (hasRoot && end === 1) {
    // return '//';
    // Backwards-compat fix:
    return '/';
  }
  return path.slice(0, end);
};

function basename(path) {
  if (typeof path !== 'string') path = path + '';

  var start = 0;
  var end = -1;
  var matchedSlash = true;
  var i;

  for (i = path.length - 1; i >= 0; --i) {
    if (path.charCodeAt(i) === 47 /*/*/) {
        // If we reached a path separator that was not part of a set of path
        // separators at the end of the string, stop now
        if (!matchedSlash) {
          start = i + 1;
          break;
        }
      } else if (end === -1) {
      // We saw the first non-path separator, mark this as the end of our
      // path component
      matchedSlash = false;
      end = i + 1;
    }
  }

  if (end === -1) return '';
  return path.slice(start, end);
}

// Uses a mixed approach for backwards-compatibility, as ext behavior changed
// in new Node.js versions, so only basename() above is backported here
exports.basename = function (path, ext) {
  var f = basename(path);
  if (ext && f.substr(-1 * ext.length) === ext) {
    f = f.substr(0, f.length - ext.length);
  }
  return f;
};

exports.extname = function (path) {
  if (typeof path !== 'string') path = path + '';
  var startDot = -1;
  var startPart = 0;
  var end = -1;
  var matchedSlash = true;
  // Track the state of characters (if any) we see before our first dot and
  // after any path separator we find
  var preDotState = 0;
  for (var i = path.length - 1; i >= 0; --i) {
    var code = path.charCodeAt(i);
    if (code === 47 /*/*/) {
        // If we reached a path separator that was not part of a set of path
        // separators at the end of the string, stop now
        if (!matchedSlash) {
          startPart = i + 1;
          break;
        }
        continue;
      }
    if (end === -1) {
      // We saw the first non-path separator, mark this as the end of our
      // extension
      matchedSlash = false;
      end = i + 1;
    }
    if (code === 46 /*.*/) {
        // If this is our first dot, mark it as the start of our extension
        if (startDot === -1)
          startDot = i;
        else if (preDotState !== 1)
          preDotState = 1;
    } else if (startDot !== -1) {
      // We saw a non-dot and non-path separator before our dot, so we should
      // have a good chance at having a non-empty extension
      preDotState = -1;
    }
  }

  if (startDot === -1 || end === -1 ||
      // We saw a non-dot character immediately before the dot
      preDotState === 0 ||
      // The (right-most) trimmed path component is exactly '..'
      preDotState === 1 && startDot === end - 1 && startDot === startPart + 1) {
    return '';
  }
  return path.slice(startDot, end);
};

function filter (xs, f) {
    if (xs.filter) return xs.filter(f);
    var res = [];
    for (var i = 0; i < xs.length; i++) {
        if (f(xs[i], i, xs)) res.push(xs[i]);
    }
    return res;
}

// String.prototype.substr - negative index don't work in IE8
var substr = 'ab'.substr(-1) === 'b'
    ? function (str, start, len) { return str.substr(start, len) }
    : function (str, start, len) {
        if (start < 0) start = str.length + start;
        return str.substr(start, len);
    }
;

/* WEBPACK VAR INJECTION */}.call(this, __webpack_require__(31)))

/***/ }),
/* 16 */
/***/ (function(module, exports) {

module.exports = require("viz.js");

/***/ }),
/* 17 */
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__.p + "voyager.worker.js";

/***/ }),
/* 18 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/colors/cyan");

/***/ }),
/* 19 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/colors/yellow");

/***/ }),
/* 20 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/Tooltip");

/***/ }),
/* 21 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/Input");

/***/ }),
/* 22 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/InputAdornment");

/***/ }),
/* 23 */
/***/ (function(module, exports) {

module.exports = require("@material-ui/core/Select");

/***/ }),
/* 24 */
/***/ (function(module, exports, __webpack_require__) {

__webpack_require__(25);
module.exports = __webpack_require__(49);


/***/ }),
/* 25 */
/***/ (function(module, exports) {

if (!Element.prototype.scrollIntoViewIfNeeded) {
    Element.prototype.scrollIntoViewIfNeeded = function (centerIfNeeded, parent) {
        centerIfNeeded = arguments.length === 0 ? true : !!centerIfNeeded;
        var parent = parent || this.parentNode, parentComputedStyle = window.getComputedStyle(parent, null), parentBorderTopWidth = parseInt(parentComputedStyle.getPropertyValue('border-top-width')), parentBorderLeftWidth = parseInt(parentComputedStyle.getPropertyValue('border-left-width')), overTop = this.offsetTop - parent.offsetTop < parent.scrollTop, overBottom = this.offsetTop - parent.offsetTop + this.clientHeight - parentBorderTopWidth >
            parent.scrollTop + parent.clientHeight, overLeft = this.offsetLeft - parent.offsetLeft < parent.scrollLeft, overRight = this.offsetLeft - parent.offsetLeft + this.clientWidth - parentBorderLeftWidth >
            parent.scrollLeft + parent.clientWidth, alignWithTop = overTop && !overBottom, hasScrolled = false;
        if ((overTop || overBottom) && centerIfNeeded) {
            parent.scrollTop =
                this.offsetTop -
                    parent.offsetTop -
                    parent.clientHeight / 2 -
                    parentBorderTopWidth +
                    this.clientHeight / 2;
            hasScrolled = true;
        }
        if ((overLeft || overRight) && centerIfNeeded) {
            parent.scrollLeft =
                this.offsetLeft -
                    parent.offsetLeft -
                    parent.clientWidth / 2 -
                    parentBorderLeftWidth +
                    this.clientWidth / 2;
            hasScrolled = true;
        }
        if ((overTop || overBottom || overLeft || overRight) && !centerIfNeeded) {
            this.scrollIntoView(alignWithTop);
            hasScrolled = true;
        }
        if (!hasScrolled &&
            parent.parentNode instanceof HTMLElement &&
            parent.clientHeight === parent.scrollHeight) {
            this.scrollIntoViewIfNeeded.call(this, centerIfNeeded, parent.parentNode);
        }
    };
}


/***/ }),
/* 26 */
/***/ (function(module, exports, __webpack_require__) {

/**
 * Modules
 */

var timestamp = __webpack_require__(27)

/**
 * Expose elapsedTime
 */

module.exports = elapsedTime

/**
 * elapsedTime
 */

function elapsedTime (start) {
  var t = start === undefined ? timestamp() : start
  return function () {
    return timestamp() - t
  }
}


/***/ }),
/* 27 */
/***/ (function(module, exports) {

/**
 * Expose timestamp
 */

module.exports = timestamp

/**
 * timestamp
 */

function timestamp () {
  return new Date().getTime()
}


/***/ }),
/* 28 */
/***/ (function(module, exports, __webpack_require__) {

/**
 * Modules
 */

var mapObj = __webpack_require__(29)

/**
 * Constants
 */

var defaultDuration = 350
var fps60 = 1000 / 60

/**
 * Expose tween
 */

module.exports = tween

/**
 * tween
 */

function tween (start, end, duration, easing, interval) {
  duration = duration === undefined ? defaultDuration : duration
  interval = interval === undefined ? fps60 : interval
  easing = easing === undefined ? linear : easing

  var frames = duration / interval

  return function (n) {
    if (frames - n < 1) return end
    return mapObj(function (val, key) {
      return tweenValue(n / frames, val, end[key], easing)
    }, start)
  }
}

/**
 * Helpers
 */

function linear (t) {
  return t
}

function tweenValue (t, start, end, ease) {
  return start + ease(t) * (end - start)
}


/***/ }),
/* 29 */
/***/ (function(module, exports) {

/**
 * Expose mapObj
 */

module.exports = map

/**
 * Map obj
 * @param  {Function} fn  map
 * @param  {Object}   obj object over which to map
 * @param  {Object}   ctx context used to map call
 * @return {Object}
 */

function map (fn, obj) {
  var result = {}
  var keys = Object.keys(obj)

  for (var i = 0, len = keys.length; i < len; ++i) {
    var key = keys[i]
    result[key] = fn.call(this, obj[key], key)
  }

  return result
}


/***/ }),
/* 30 */
/***/ (function(module, exports) {

/**
 * Constants
 */

var rafInterval = 1000 / 60  // 60 frames per second

/**
 * Expose raf and cancel
 */

if (typeof window === 'undefined' || !window.requestAnimationFrame) {
  exports = module.exports = polyfill
  exports.cancel = clearTimeout
} else {
  exports = module.exports = requestAnimationFrame.bind(window)
  exports.cancel = window.cancelAnimationFrame.bind(window)
}

/**
 * Polyfill
 */

var prev = new Date().getTime()

function polyfill (fn) {
  var cur = new Date().getTime()
  var ms = Math.max(0, rafInterval - (cur - prev))
  prev = cur
  return setTimeout(fn, ms)
}


/***/ }),
/* 31 */
/***/ (function(module, exports) {

// shim for using process in browser
var process = module.exports = {};

// cached from whatever global is present so that test runners that stub it
// don't break things.  But we need to wrap it in a try catch in case it is
// wrapped in strict mode code which doesn't define any globals.  It's inside a
// function because try/catches deoptimize in certain engines.

var cachedSetTimeout;
var cachedClearTimeout;

function defaultSetTimout() {
    throw new Error('setTimeout has not been defined');
}
function defaultClearTimeout () {
    throw new Error('clearTimeout has not been defined');
}
(function () {
    try {
        if (typeof setTimeout === 'function') {
            cachedSetTimeout = setTimeout;
        } else {
            cachedSetTimeout = defaultSetTimout;
        }
    } catch (e) {
        cachedSetTimeout = defaultSetTimout;
    }
    try {
        if (typeof clearTimeout === 'function') {
            cachedClearTimeout = clearTimeout;
        } else {
            cachedClearTimeout = defaultClearTimeout;
        }
    } catch (e) {
        cachedClearTimeout = defaultClearTimeout;
    }
} ())
function runTimeout(fun) {
    if (cachedSetTimeout === setTimeout) {
        //normal enviroments in sane situations
        return setTimeout(fun, 0);
    }
    // if setTimeout wasn't available but was latter defined
    if ((cachedSetTimeout === defaultSetTimout || !cachedSetTimeout) && setTimeout) {
        cachedSetTimeout = setTimeout;
        return setTimeout(fun, 0);
    }
    try {
        // when when somebody has screwed with setTimeout but no I.E. maddness
        return cachedSetTimeout(fun, 0);
    } catch(e){
        try {
            // When we are in I.E. but the script has been evaled so I.E. doesn't trust the global object when called normally
            return cachedSetTimeout.call(null, fun, 0);
        } catch(e){
            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error
            return cachedSetTimeout.call(this, fun, 0);
        }
    }


}
function runClearTimeout(marker) {
    if (cachedClearTimeout === clearTimeout) {
        //normal enviroments in sane situations
        return clearTimeout(marker);
    }
    // if clearTimeout wasn't available but was latter defined
    if ((cachedClearTimeout === defaultClearTimeout || !cachedClearTimeout) && clearTimeout) {
        cachedClearTimeout = clearTimeout;
        return clearTimeout(marker);
    }
    try {
        // when when somebody has screwed with setTimeout but no I.E. maddness
        return cachedClearTimeout(marker);
    } catch (e){
        try {
            // When we are in I.E. but the script has been evaled so I.E. doesn't  trust the global object when called normally
            return cachedClearTimeout.call(null, marker);
        } catch (e){
            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error.
            // Some versions of I.E. have different rules for clearTimeout vs setTimeout
            return cachedClearTimeout.call(this, marker);
        }
    }



}
var queue = [];
var draining = false;
var currentQueue;
var queueIndex = -1;

function cleanUpNextTick() {
    if (!draining || !currentQueue) {
        return;
    }
    draining = false;
    if (currentQueue.length) {
        queue = currentQueue.concat(queue);
    } else {
        queueIndex = -1;
    }
    if (queue.length) {
        drainQueue();
    }
}

function drainQueue() {
    if (draining) {
        return;
    }
    var timeout = runTimeout(cleanUpNextTick);
    draining = true;

    var len = queue.length;
    while(len) {
        currentQueue = queue;
        queue = [];
        while (++queueIndex < len) {
            if (currentQueue) {
                currentQueue[queueIndex].run();
            }
        }
        queueIndex = -1;
        len = queue.length;
    }
    currentQueue = null;
    draining = false;
    runClearTimeout(timeout);
}

process.nextTick = function (fun) {
    var args = new Array(arguments.length - 1);
    if (arguments.length > 1) {
        for (var i = 1; i < arguments.length; i++) {
            args[i - 1] = arguments[i];
        }
    }
    queue.push(new Item(fun, args));
    if (queue.length === 1 && !draining) {
        runTimeout(drainQueue);
    }
};

// v8 likes predictible objects
function Item(fun, array) {
    this.fun = fun;
    this.array = array;
}
Item.prototype.run = function () {
    this.fun.apply(null, this.array);
};
process.title = 'browser';
process.browser = true;
process.env = {};
process.argv = [];
process.version = ''; // empty string to avoid regexp issues
process.versions = {};

function noop() {}

process.on = noop;
process.addListener = noop;
process.once = noop;
process.off = noop;
process.removeListener = noop;
process.removeAllListeners = noop;
process.emit = noop;
process.prependListener = noop;
process.prependOnceListener = noop;

process.listeners = function (name) { return [] }

process.binding = function (name) {
    throw new Error('process.binding is not supported');
};

process.cwd = function () { return '/' };
process.chdir = function (dir) {
    throw new Error('process.chdir is not supported');
};
process.umask = function() { return 0; };


/***/ }),
/* 32 */
/***/ (function(module, exports) {

module.exports = "<symbol viewBox=\"0 0 600 600\" id=\"RelayIcon\"><g fill=\"#F26B00\"><path d=\"M142.536 198.858c0 26.36-21.368 47.72-47.72 47.72-26.36 0-47.722-21.36-47.722-47.72s21.36-47.72 47.72-47.72c26.355 0 47.722 21.36 47.722 47.72\"/><path d=\"M505.18 414.225H238.124c-35.25 0-63.926-28.674-63.926-63.923s28.678-63.926 63.926-63.926h120.78c20.816 0 37.753-16.938 37.753-37.756s-16.938-37.756-37.753-37.756H94.81c-7.227 0-13.086-5.86-13.086-13.085 0-7.227 5.86-13.086 13.085-13.086h264.093c35.25 0 63.923 28.678 63.923 63.926s-28.674 63.923-63.923 63.923h-120.78c-20.82 0-37.756 16.938-37.756 37.76 0 20.816 16.938 37.753 37.756 37.753H505.18c7.227 0 13.086 5.86 13.086 13.085 0 7.226-5.858 13.085-13.085 13.085z\"/><path d=\"M457.464 401.142c0-26.36 21.36-47.72 47.72-47.72s47.72 21.36 47.72 47.72-21.36 47.72-47.72 47.72-47.72-21.36-47.72-47.72\"/></g></symbol>"

/***/ }),
/* 33 */
/***/ (function(module, exports) {

module.exports = "<symbol viewBox=\"-50 -50 580 700\" id=\"DeprecatedIcon\">\n    <path style=\"fill:#f74036;\" d=\"M481.249,420.625l-219-380c-4-7-11-11-19-11s-15,4-19,11l-221,382c-10,17,5,33,19,33h442\n        c12,0,21-10,21-22C485.249,428.625,484.249,424.625,481.249,420.625z M60.249,411.625l183-317l183,317H60.249z\"/>\n    <path style=\"fill:#f74036;\" d=\"M221.249,192.625v97c0,12,10,22,22,22s22-10,22-22v-97c0-12-10-22-22-22\n        S221.249,180.625,221.249,192.625z\"/>\n    <path style=\"fill:#f74036;\" d=\"M243.249,330.625c-15,0-27,12-27,27s12,28,27,28s27-13,27-28S258.249,330.625,243.249,330.625z\"/>\n</symbol>"

/***/ }),
/* 34 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 35 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 36 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 37 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 38 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 39 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 40 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 41 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 42 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 43 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 44 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 45 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 46 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 47 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 48 */
/***/ (function(module, exports) {

// removed by extract-text-webpack-plugin

/***/ }),
/* 49 */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);

// EXTERNAL MODULE: external "react"
var external_react_ = __webpack_require__(0);
var external_react_default = /*#__PURE__*/__webpack_require__.n(external_react_);

// EXTERNAL MODULE: external "react-dom"
var external_react_dom_ = __webpack_require__(12);

// EXTERNAL MODULE: external "graphql/utilities"
var utilities_ = __webpack_require__(13);

// EXTERNAL MODULE: external "lodash"
var external_lodash_ = __webpack_require__(1);

// EXTERNAL MODULE: external "graphql"
var external_graphql_ = __webpack_require__(8);

// CONCATENATED MODULE: ./src/introspection/utils.ts

function stringifyWrappers(wrappers) {
    var left = '';
    var right = '';
    for (var _i = 0, wrappers_1 = wrappers; _i < wrappers_1.length; _i++) {
        var wrapper = wrappers_1[_i];
        switch (wrapper) {
            case 'NON_NULL':
                right = '!' + right;
                break;
            case 'LIST':
                left = left + '[';
                right = ']' + right;
                break;
        }
    }
    return [left, right];
}
function buildId() {
    var parts = [];
    for (var _i = 0; _i < arguments.length; _i++) {
        parts[_i] = arguments[_i];
    }
    return parts.join('::');
}
function typeNameToId(name) {
    return buildId('TYPE', name);
}
function extractTypeId(id) {
    var _a = id.split('::'), type = _a[1];
    return buildId('TYPE', type);
}
function isSystemType(type) {
    return external_lodash_["startsWith"](type.name, '__');
}
function isBuiltInScalarType(type) {
    return ['Int', 'Float', 'String', 'Boolean', 'ID'].indexOf(type.name) !== -1;
}
function isScalarType(type) {
    return type.kind === 'SCALAR' || type.kind === 'ENUM';
}
function isObjectType(type) {
    return type.kind === 'OBJECT';
}
function isInputObjectType(type) {
    return type.kind === 'INPUT_OBJECT';
}

// CONCATENATED MODULE: ./src/introspection/introspection.ts



function unwrapType(type, wrappers) {
    while (type.kind === 'NON_NULL' || type.kind == 'LIST') {
        wrappers.push(type.kind);
        type = type.ofType;
    }
    return type.name;
}
function convertArg(inArg) {
    var outArg = {
        name: inArg.name,
        description: inArg.description,
        defaultValue: inArg.defaultValue,
        typeWrappers: [],
    };
    outArg.type = unwrapType(inArg.type, outArg.typeWrappers);
    return outArg;
}
var convertInputField = convertArg;
function convertField(inField) {
    var outField = {
        name: inField.name,
        description: inField.description,
        typeWrappers: [],
        isDeprecated: inField.isDeprecated,
    };
    outField.type = unwrapType(inField.type, outField.typeWrappers);
    outField.args = external_lodash_(inField.args)
        .map(convertArg)
        .keyBy('name')
        .value();
    if (outField.isDeprecated)
        outField.deprecationReason = inField.deprecationReason;
    return outField;
}
function convertType(inType) {
    var outType = {
        kind: inType.kind,
        name: inType.name,
        description: inType.description,
    };
    switch (inType.kind) {
        case 'OBJECT':
            outType.interfaces = external_lodash_(inType.interfaces)
                .map('name')
                .uniq()
                .value();
            outType.fields = external_lodash_(inType.fields)
                .map(convertField)
                .keyBy('name')
                .value();
            break;
        case 'INTERFACE':
            outType.derivedTypes = external_lodash_(inType.possibleTypes)
                .map('name')
                .uniq()
                .value();
            outType.fields = external_lodash_(inType.fields)
                .map(convertField)
                .keyBy('name')
                .value();
            break;
        case 'UNION':
            outType.possibleTypes = external_lodash_(inType.possibleTypes)
                .map('name')
                .uniq()
                .value();
            break;
        case 'ENUM':
            outType.enumValues = inType.enumValues.slice();
            break;
        case 'INPUT_OBJECT':
            outType.inputFields = external_lodash_(inType.inputFields)
                .map(convertInputField)
                .keyBy('name')
                .value();
            break;
    }
    return outType;
}
function simplifySchema(inSchema) {
    return {
        types: external_lodash_(inSchema.types)
            .map(convertType)
            .keyBy('name')
            .value(),
        queryType: inSchema.queryType.name,
        mutationType: external_lodash_["get"](inSchema, 'mutationType.name', null),
        subscriptionType: external_lodash_["get"](inSchema, 'subscriptionType.name', null),
    };
}
function markRelayTypes(schema) {
    var nodeType = schema.types[typeNameToId('Node')];
    if (nodeType)
        nodeType.isRelayType = true;
    var pageInfoType = schema.types[typeNameToId('PageInfo')];
    if (pageInfoType)
        pageInfoType.isRelayType = true;
    var edgeTypesMap = {};
    external_lodash_["each"](schema.types, function (type) {
        if (!external_lodash_["isEmpty"](type.interfaces)) {
            type.interfaces = external_lodash_["reject"](type.interfaces, function (baseType) { return baseType.type.name === 'Node'; });
        }
        external_lodash_["each"](type.fields, function (field) {
            var connectionType = field.type;
            if (!/.Connection$/.test(connectionType.name) ||
                connectionType.kind !== 'OBJECT' ||
                !connectionType.fields.edges) {
                return;
            }
            var edgesType = connectionType.fields.edges.type;
            if (edgesType.kind !== 'OBJECT' || !edgesType.fields.node) {
                return;
            }
            var nodeType = edgesType.fields.node.type;
            connectionType.isRelayType = true;
            edgesType.isRelayType = true;
            edgeTypesMap[edgesType.name] = nodeType;
            field.relayType = field.type;
            field.type = nodeType;
            field.typeWrappers = ['LIST'];
            var relayArgNames = ['first', 'last', 'before', 'after'];
            var isRelayArg = function (arg) { return relayArgNames.includes(arg.name); };
            field.relayArgs = external_lodash_["pickBy"](field.args, isRelayArg);
            field.args = external_lodash_["omitBy"](field.args, isRelayArg);
        });
    });
    external_lodash_["each"](schema.types, function (type) {
        external_lodash_["each"](type.fields, function (field) {
            var realType = edgeTypesMap[field.type.name];
            if (realType === undefined)
                return;
            field.relayType = field.type;
            field.type = realType;
        });
    });
    var queryType = schema.queryType;
    var query = schema.types[queryType.id];
    if (external_lodash_["get"](query, 'fields.node.type.isRelayType')) {
        delete query.fields['node'];
    }
    //GitHub use `nodes` instead of `node`.
    if (external_lodash_["get"](query, 'fields.nodes.type.isRelayType')) {
        delete query.fields['nodes'];
    }
    if (external_lodash_["get"](query, 'fields.relay.type') === queryType) {
        delete query.fields['relay'];
    }
}
function markDeprecated(schema) {
    // Remove deprecated fields.
    external_lodash_["each"](schema.types, function (type) {
        type.fields = external_lodash_["pickBy"](type.fields, function (field) { return !field.isDeprecated; });
    });
    // We can't remove types that end up being empty
    // because we cannot be sure that the @deprecated directives where
    // consistently added to the schema we're handling.
    //
    // Entities may have non deprecated fields pointing towards entities
    // which are deprecated.
}
function assignTypesAndIDs(schema) {
    schema.queryType = schema.types[schema.queryType];
    schema.mutationType = schema.types[schema.mutationType];
    schema.subscriptionType = schema.types[schema.subscriptionType];
    external_lodash_["each"](schema.types, function (type) {
        type.id = typeNameToId(type.name);
        external_lodash_["each"](type.inputFields, function (field) {
            field.id = "FIELD::" + type.name + "::" + field.name;
            field.type = schema.types[field.type];
        });
        external_lodash_["each"](type.fields, function (field) {
            field.id = "FIELD::" + type.name + "::" + field.name;
            field.type = schema.types[field.type];
            external_lodash_["each"](field.args, function (arg) {
                arg.id = "ARGUMENT::" + type.name + "::" + field.name + "::" + arg.name;
                arg.type = schema.types[arg.type];
            });
        });
        if (!external_lodash_["isEmpty"](type.possibleTypes)) {
            type.possibleTypes = external_lodash_["map"](type.possibleTypes, function (possibleType) { return ({
                id: "POSSIBLE_TYPE::" + type.name + "::" + possibleType,
                type: schema.types[possibleType],
            }); });
        }
        if (!external_lodash_["isEmpty"](type.derivedTypes)) {
            type.derivedTypes = external_lodash_["map"](type.derivedTypes, function (derivedType) { return ({
                id: "DERIVED_TYPE::" + type.name + "::" + derivedType,
                type: schema.types[derivedType],
            }); });
        }
        if (!external_lodash_["isEmpty"](type.interfaces)) {
            type.interfaces = external_lodash_["map"](type.interfaces, function (baseType) { return ({
                id: "INTERFACE::" + type.name + "::" + baseType,
                type: schema.types[baseType],
            }); });
        }
    });
    schema.types = external_lodash_["keyBy"](schema.types, 'id');
}
function getSchema(introspection, sortByAlphabet, skipRelay, skipDeprecated) {
    if (!introspection)
        return null;
    var schema = Object(external_graphql_["buildClientSchema"])(introspection.data);
    if (sortByAlphabet) {
        schema = Object(external_graphql_["lexicographicSortSchema"])(schema);
    }
    introspection = Object(external_graphql_["introspectionFromSchema"])(schema, { descriptions: true });
    var simpleSchema = simplifySchema(introspection.__schema);
    assignTypesAndIDs(simpleSchema);
    if (skipRelay) {
        markRelayTypes(simpleSchema);
    }
    if (skipDeprecated) {
        markDeprecated(simpleSchema);
    }
    return simpleSchema;
}

// CONCATENATED MODULE: ./src/introspection/index.ts



// CONCATENATED MODULE: ./src/graph/type-graph.ts


function type_graph_isNode(type) {
    return !(isScalarType(type) || isInputObjectType(type) || isSystemType(type) || type.isRelayType);
}
function getDefaultRoot(schema) {
    return schema.queryType.name;
}
function getTypeGraph(schema, rootType, hideRoot) {
    if (schema === null)
        return null;
    var rootId = typeNameToId(rootType || getDefaultRoot(schema));
    return buildGraph(rootId);
    function getEdgeTargets(type) {
        return external_lodash_(external_lodash_["values"](type.fields).concat((type.derivedTypes || []), (type.possibleTypes || [])))
            .map('type')
            .filter(type_graph_isNode)
            .map('id')
            .value();
    }
    function buildGraph(rootId) {
        var typeIds = [rootId];
        var nodes = [];
        var types = external_lodash_["keyBy"](schema.types, 'id');
        for (var i = 0; i < typeIds.length; ++i) {
            var id = typeIds[i];
            if (typeIds.indexOf(id) < i)
                continue;
            var type = types[id];
            nodes.push(type);
            typeIds.push.apply(typeIds, getEdgeTargets(type));
        }
        return {
            rootId: rootId,
            nodes: hideRoot ? external_lodash_["omit"](external_lodash_["keyBy"](nodes, 'id'), [rootId]) : external_lodash_["keyBy"](nodes, 'id'),
        };
    }
}

// EXTERNAL MODULE: external "svg-pan-zoom"
var external_svg_pan_zoom_ = __webpack_require__(14);

// EXTERNAL MODULE: ./node_modules/@f/animate/lib/index.js
var lib = __webpack_require__(9);

// EXTERNAL MODULE: ./node_modules/path-browserify/index.js
var path_browserify = __webpack_require__(15);

// CONCATENATED MODULE: ./src/utils/dom-helpers.ts
function forEachNode(parent, selector, fn) {
    var $nodes = parent.querySelectorAll(selector);
    for (var i = 0; i < $nodes.length; i++) {
        fn($nodes[i]);
    }
}
function addClass(parent, selector, className) {
    forEachNode(parent, selector, function (node) { return node.classList.add(className); });
}
function removeClass(parent, selector, className) {
    forEachNode(parent, selector, function (node) { return node.classList.remove(className); });
}
function stringToSvg(svgString) {
    var svgDoc = new DOMParser().parseFromString(svgString, 'image/svg+xml');
    return document.importNode(svgDoc.documentElement, true);
}

// CONCATENATED MODULE: ./src/utils/highlight.tsx


function highlightTerm(content, term) {
    if (!term) {
        return content;
    }
    var re = new RegExp('(' + external_lodash_["escapeRegExp"](term) + ')', 'gi');
    var result = content.split(re);
    // Apply highlight to all odd elements
    for (var i = 1, length = result.length; i < length; i += 2) {
        result[i] = (external_react_["createElement"]("span", { key: i, style: { backgroundColor: '#ffff03' } }, result[i]));
    }
    return result;
}

// CONCATENATED MODULE: ./src/utils/index.ts

// similar to node __dirname
var utils_dirname;


function isMatch(sourceText, searchValue) {
    if (!searchValue) {
        return true;
    }
    try {
        var escaped = searchValue.replace(/[^_0-9A-Za-z]/g, function (ch) { return '\\' + ch; });
        return sourceText.search(new RegExp(escaped, 'i')) !== -1;
    }
    catch (e) {
        return sourceText.toLowerCase().indexOf(searchValue.toLowerCase()) !== -1;
    }
}
function utils_loadWorker(path, relative) {
    var url = relative ? utils_dirname + '/' + path : path;
    return fetch(url)
        .then(function (response) { return response.text(); })
        .then(function (payload) {
        // HACK: to increase viz.js memory size from 16mb to 256mb
        // should use response.blob()
        payload = payload
            .replace('||16777216;', '||(16777216 * 16);')
            .replace('||5242880;', '||(5242880 * 16);');
        var script = new Blob([payload], { type: 'application/javascript' });
        var url = URL.createObjectURL(script);
        return new Worker(url);
    });
}
/*
  get current script URL
*/
function getJsUrl() {
    var id = +new Date() + Math.random();
    try {
        // write empty script to the document. It will get placed directly after the current script
        document.write("<script id=\"dummy" + id + "\"></script>");
        // find appended script and return src of the previous script which is the current script
        return document.getElementById('dummy' + id).previousSibling.src;
    }
    catch (e) {
        return '';
    }
}
utils_dirname = path_browserify["dirname"](getJsUrl());

// CONCATENATED MODULE: ./src/graph/viewport.ts





var viewport_Viewport = /** @class */ (function () {
    function Viewport(svgString, container, onSelectNode, onSelectEdge) {
        var _this = this;
        this.container = container;
        this.onSelectNode = onSelectNode;
        this.onSelectEdge = onSelectEdge;
        this.container.innerHTML = '';
        this.$svg = stringToSvg(svgString);
        this.container.appendChild(this.$svg);
        // Allow the SVG dimensions to be computed
        // Quick fix for SVG manipulation issues.
        setTimeout(function () { return _this.enableZoom(); }, 0);
        this.bindClick();
        this.bindHover();
        this.resize();
        window.addEventListener('resize', function () { return _this.resize(); });
    }
    Viewport.prototype.resize = function () {
        var bbRect = this.container.getBoundingClientRect();
        this.offsetLeft = bbRect.left;
        this.offsetTop = bbRect.top;
        if (this.zoomer !== undefined) {
            this.zoomer.resize();
        }
    };
    Viewport.prototype.enableZoom = function () {
        var svgHeight = this.$svg['height'].baseVal.value;
        var svgWidth = this.$svg['width'].baseVal.value;
        var bbRect = this.container.getBoundingClientRect();
        this.maxZoom = Math.max(svgHeight / bbRect.height, svgWidth / bbRect.width);
        this.zoomer = external_svg_pan_zoom_(this.$svg, {
            zoomScaleSensitivity: 0.25,
            minZoom: 0.95,
            maxZoom: this.maxZoom,
            controlIconsEnabled: true,
        });
        this.zoomer.zoom(0.95);
    };
    Viewport.prototype.bindClick = function () {
        var _this = this;
        var dragged = false;
        var moveHandler = function () { return (dragged = true); };
        this.$svg.addEventListener('mousedown', function () {
            dragged = false;
            setTimeout(function () { return _this.$svg.addEventListener('mousemove', moveHandler); });
        });
        this.$svg.addEventListener('mouseup', function (event) {
            _this.$svg.removeEventListener('mousemove', moveHandler);
            if (dragged)
                return;
            var target = event.target;
            if (isLink(target)) {
                var typeId = typeNameToId(target.textContent);
                _this.focusElement(typeId);
            }
            else if (viewport_isNode(target)) {
                var $node = getParent(target, 'node');
                _this.onSelectNode($node.id);
            }
            else if (isEdge(target)) {
                var $edge = getParent(target, 'edge');
                _this.onSelectEdge(edgeSource($edge).id);
            }
            else if (!isControl(target)) {
                _this.onSelectNode(null);
            }
        });
    };
    Viewport.prototype.bindHover = function () {
        var $prevHovered = null;
        var $prevHoveredEdge = null;
        function clearSelection() {
            if ($prevHovered)
                $prevHovered.classList.remove('hovered');
            if ($prevHoveredEdge)
                $prevHoveredEdge.classList.remove('hovered');
        }
        this.$svg.addEventListener('mousemove', function (event) {
            var target = event.target;
            if (isEdgeSource(target)) {
                var $sourceGroup = getParent(target, 'edge-source');
                if ($sourceGroup.classList.contains('hovered'))
                    return;
                clearSelection();
                $sourceGroup.classList.add('hovered');
                $prevHovered = $sourceGroup;
                var $edge = edgeFrom($sourceGroup.id);
                $edge.classList.add('hovered');
                $prevHoveredEdge = $edge;
            }
            else {
                clearSelection();
            }
        });
    };
    Viewport.prototype.selectNodeById = function (id) {
        this.deselectNode();
        if (id === null) {
            this.$svg.classList.remove('selection-active');
            return;
        }
        this.$svg.classList.add('selection-active');
        var $selected = document.getElementById(id);
        this.selectNode($selected);
    };
    Viewport.prototype.selectNode = function (node) {
        node.classList.add('selected');
        external_lodash_["each"](edgesFromNode(node), function ($edge) {
            $edge.classList.add('highlighted');
            edgeTarget($edge).classList.add('selected-reachable');
        });
        external_lodash_["each"](edgesTo(node.id), function ($edge) {
            $edge.classList.add('highlighted');
            edgeSource($edge).parentElement.classList.add('selected-reachable');
        });
    };
    Viewport.prototype.selectEdgeById = function (id) {
        removeClass(this.$svg, '.edge.selected', 'selected');
        removeClass(this.$svg, '.edge-source.selected', 'selected');
        removeClass(this.$svg, '.field.selected', 'selected');
        if (id === null)
            return;
        var $selected = document.getElementById(id);
        if ($selected) {
            var $edge = edgeFrom($selected.id);
            if ($edge)
                $edge.classList.add('selected');
            $selected.classList.add('selected');
        }
    };
    Viewport.prototype.deselectNode = function () {
        removeClass(this.$svg, '.node.selected', 'selected');
        removeClass(this.$svg, '.highlighted', 'highlighted');
        removeClass(this.$svg, '.selected-reachable', 'selected-reachable');
    };
    Viewport.prototype.focusElement = function (id) {
        var bbBox = document.getElementById(id).getBoundingClientRect();
        var currentPan = this.zoomer.getPan();
        var viewPortSizes = this.zoomer.getSizes();
        currentPan.x += viewPortSizes.width / 2 - bbBox.width / 2;
        currentPan.y += viewPortSizes.height / 2 - bbBox.height / 2;
        var zoomUpdateToFit = 1.2 * Math.max(bbBox.height / viewPortSizes.height, bbBox.width / viewPortSizes.width);
        var newZoom = this.zoomer.getZoom() / zoomUpdateToFit;
        var recomendedZoom = this.maxZoom * 0.6;
        if (newZoom > recomendedZoom)
            newZoom = recomendedZoom;
        var newX = currentPan.x - bbBox.left + this.offsetLeft;
        var newY = currentPan.y - bbBox.top + this.offsetTop;
        this.animatePanAndZoom(newX, newY, newZoom);
    };
    Viewport.prototype.animatePanAndZoom = function (x, y, zoomEnd) {
        var _this = this;
        var pan = this.zoomer.getPan();
        var panEnd = { x: x, y: y };
        lib(pan, panEnd, function (props) {
            _this.zoomer.pan({ x: props.x, y: props.y });
            if (props === panEnd) {
                var zoom = _this.zoomer.getZoom();
                lib({ zoom: zoom }, { zoom: zoomEnd }, function (props) {
                    _this.zoomer.zoom(props.zoom);
                });
            }
        });
    };
    Viewport.prototype.destroy = function () {
        window.removeEventListener('resize', this.resize);
        try {
            this.zoomer.destroy();
        }
        catch (e) {
            // skip
        }
    };
    return Viewport;
}());

function getParent(elem, className) {
    while (elem && elem.tagName !== 'svg') {
        if (elem.classList.contains(className))
            return elem;
        elem = elem.parentNode;
    }
    return null;
}
function viewport_isNode(elem) {
    return getParent(elem, 'node') != null;
}
function isEdge(elem) {
    return getParent(elem, 'edge') != null;
}
function isLink(elem) {
    return elem.classList.contains('type-link');
}
function isEdgeSource(elem) {
    return getParent(elem, 'edge-source') != null;
}
function isControl(elem) {
    if (!(elem instanceof SVGElement))
        return false;
    return elem.className.baseVal.startsWith('svg-pan-zoom');
}
function edgeSource(edge) {
    return document.getElementById(edge['dataset']['from']);
}
function edgeTarget(edge) {
    return document.getElementById(edge['dataset']['to']);
}
function edgeFrom(id) {
    return document.querySelector(".edge[data-from='" + id + "']");
}
function edgesFromNode($node) {
    var edges = [];
    forEachNode($node, '.edge-source', function ($source) {
        var $edge = edgeFrom($source.id);
        edges.push($edge);
    });
    return edges;
}
function edgesTo(id) {
    return external_lodash_["toArray"](document.querySelectorAll(".edge[data-to='" + id + "']"));
}

// CONCATENATED MODULE: ./src/graph/dot.ts


function getDot(typeGraph, displayOptions) {
    function isNode(type) {
        return typeGraph.nodes[type.id] !== undefined;
    }
    return (typeGraph &&
        "\n    digraph {\n      graph [\n        rankdir = \"LR\"\n      ];\n      node [\n        fontsize = \"16\"\n        fontname = \"helvetica, open-sans\"\n        shape = \"plaintext\"\n      ];\n      edge [\n      ];\n      ranksep = 2.0\n      " + objectValues(typeGraph.nodes, function (node) { return "\n        \"" + node.name + "\" [\n          id = \"" + node.id + "\"\n          label = " + nodeLabel(node) + "\n        ]\n        " + objectValues(node.fields, function (field) {
            return isNode(field.type)
                ? "\n          \"" + node.name + "\":\"" + field.name + "\" -> \"" + field.type.name + "\" [\n            id = \"" + field.id + " => " + field.type.id + "\"\n            label = \"" + node.name + ":" + field.name + "\"\n          ]\n        "
                : '';
        }) + ";\n        " + array(node.possibleTypes, function (_a) {
            var id = _a.id, type = _a.type;
            return "\n          \"" + node.name + "\":\"" + type.name + "\" -> \"" + type.name + "\" [\n            id = \"" + id + " => " + type.id + "\"\n            style = \"dashed\"\n          ]\n        ";
        }) + "\n        " + array(node.derivedTypes, function (_a) {
            var id = _a.id, type = _a.type;
            return "\n          \"" + node.name + "\":\"" + type.name + "\" -> \"" + type.name + "\" [\n            id = \"" + id + " => " + type.id + "\"\n            style = \"dotted\"\n          ]\n        ";
        }) + "\n      "; }) + "\n    }\n  ");
    function nodeLabel(node) {
        var htmlID = HtmlId('TYPE_TITLE::' + node.name);
        var kindLabel = node.kind !== 'OBJECT' ? '&lt;&lt;' + node.kind.toLowerCase() + '&gt;&gt;' : '';
        return "\n      <<TABLE ALIGN=\"LEFT\" BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"5\">\n        <TR>\n          <TD CELLPADDING=\"4\" " + htmlID + "><FONT POINT-SIZE=\"18\">" + node.name + "</FONT><BR/>" + kindLabel + "</TD>\n        </TR>\n        " + objectValues(node.fields, nodeField) + "\n        " + dot_possibleTypes(node) + "\n        " + dot_derivedTypes(node) + "\n      </TABLE>>\n    ";
    }
    function canDisplayRow(type) {
        if (type.kind === 'SCALAR' || type.kind === 'ENUM') {
            return displayOptions.showLeafFields;
        }
        return true;
    }
    function nodeField(field) {
        var relayIcon = field.relayType ? TEXT('{R}') : '';
        var deprecatedIcon = field.isDeprecated ? TEXT('{D}') : '';
        var parts = stringifyWrappers(field.typeWrappers).map(TEXT);
        return canDisplayRow(field.type)
            ? "\n      <TR>\n        <TD " + HtmlId(field.id) + " ALIGN=\"LEFT\" PORT=\"" + field.name + "\">\n          <TABLE CELLPADDING=\"0\" CELLSPACING=\"0\" BORDER=\"0\">\n            <TR>\n              <TD ALIGN=\"LEFT\">" + field.name + "<FONT>  </FONT></TD>\n              <TD ALIGN=\"RIGHT\">" + deprecatedIcon + relayIcon + parts[0] + field.type.name + parts[1] + "</TD>\n            </TR>\n          </TABLE>\n        </TD>\n      </TR>\n    "
            : '';
    }
}
function dot_possibleTypes(node) {
    var possibleTypes = node.possibleTypes;
    if (external_lodash_["isEmpty"](possibleTypes)) {
        return '';
    }
    return "\n    <TR>\n      <TD>possible types</TD>\n    </TR>\n    " + array(possibleTypes, function (_a) {
        var id = _a.id, type = _a.type;
        return "\n      <TR>\n        <TD " + HtmlId(id) + " ALIGN=\"LEFT\" PORT=\"" + type.name + "\">" + type.name + "</TD>\n      </TR>\n    ";
    }) + "\n  ";
}
function dot_derivedTypes(node) {
    var derivedTypes = node.derivedTypes;
    if (external_lodash_["isEmpty"](derivedTypes)) {
        return '';
    }
    return "\n    <TR>\n      <TD>implementations</TD>\n    </TR>\n    " + array(derivedTypes, function (_a) {
        var id = _a.id, type = _a.type;
        return "\n      <TR>\n        <TD " + HtmlId(id) + " ALIGN=\"LEFT\" PORT=\"" + type.name + "\">" + type.name + "</TD>\n      </TR>\n    ";
    }) + "\n  ";
}
function objectValues(object, stringify) {
    return external_lodash_["values"](object)
        .map(stringify)
        .join('\n');
}
function array(array, stringify) {
    return array ? array.map(stringify).join('\n') : '';
}
function HtmlId(id) {
    return 'HREF="remove_me_url" ID="' + id + '"';
}
function TEXT(str) {
    if (str === '')
        return '';
    str = str.replace(/]/, '&#93;');
    return '<FONT>' + str + '</FONT>';
}

// EXTERNAL MODULE: external "viz.js"
var external_viz_js_ = __webpack_require__(16);
var external_viz_js_default = /*#__PURE__*/__webpack_require__.n(external_viz_js_);

// EXTERNAL MODULE: ./node_modules/viz.js/full.render.js
var full_render = __webpack_require__(17);
var full_render_default = /*#__PURE__*/__webpack_require__.n(full_render);

// CONCATENATED MODULE: ./src/graph/svg-renderer.ts





var RelayIconSvg = __webpack_require__(32);
var DeprecatedIconSvg = __webpack_require__(33);
var svgns = 'http://www.w3.org/2000/svg';
var xlinkns = 'http://www.w3.org/1999/xlink';
var svg_renderer_SVGRender = /** @class */ (function () {
    function SVGRender(workerURI, loadWorker) {
        if (loadWorker === void 0) { loadWorker = utils_loadWorker; }
        this.vizPromise = loadWorker(workerURI || full_render_default.a, !workerURI).then(function (worker) { return new external_viz_js_default.a({ worker: worker }); });
    }
    SVGRender.prototype.renderSvg = function (typeGraph, displayOptions) {
        return this.vizPromise
            .then(function (viz) {
            console.time('Rendering Graph');
            var dot = getDot(typeGraph, displayOptions);
            return viz.renderString(dot);
        })
            .then(function (rawSVG) {
            var svg = preprocessVizSVG(rawSVG);
            console.timeEnd('Rendering Graph');
            return svg;
        });
    };
    return SVGRender;
}());

function preprocessVizSVG(svgString) {
    //Add Relay and Deprecated icons
    svgString = svgString.replace(/<svg [^>]*>/, '$&' + RelayIconSvg);
    svgString = svgString.replace(/<svg [^>]*>/, '$&' + DeprecatedIconSvg);
    var svg = stringToSvg(svgString);
    forEachNode(svg, 'a', function ($a) {
        var $g = $a.parentNode;
        var $docFrag = document.createDocumentFragment();
        while ($a.firstChild) {
            var $child = $a.firstChild;
            $docFrag.appendChild($child);
        }
        $g.replaceChild($docFrag, $a);
        $g.id = $g.id.replace(/^a_/, '');
    });
    forEachNode(svg, 'title', function ($el) { return $el.remove(); });
    var edgesSources = {};
    forEachNode(svg, '.edge', function ($edge) {
        var _a = $edge.id.split(' => '), from = _a[0], to = _a[1];
        $edge.removeAttribute('id');
        $edge.setAttribute('data-from', from);
        $edge.setAttribute('data-to', to);
        edgesSources[from] = true;
    });
    forEachNode(svg, '[id]', function ($el) {
        var _a = $el.id.split('::'), tag = _a[0], restOfId = _a.slice(1);
        if (external_lodash_["size"](restOfId) < 1)
            return;
        $el.classList.add(tag.toLowerCase().replace(/_/, '-'));
    });
    forEachNode(svg, 'g.edge path', function ($path) {
        var $newPath = $path.cloneNode();
        $newPath.classList.add('hover-path');
        $newPath.removeAttribute('stroke-dasharray');
        $path.parentNode.appendChild($newPath);
    });
    forEachNode(svg, '.field', function ($field) {
        var texts = $field.querySelectorAll('text');
        texts[0].classList.add('field-name');
        //Remove spaces used for text alligment
        texts[1].remove();
        if (edgesSources[$field.id])
            $field.classList.add('edge-source');
        for (var i = 2; i < texts.length; ++i) {
            var str = texts[i].innerHTML;
            if (str === '{R}' || str == '{D}') {
                var $iconPlaceholder = texts[i];
                var height = 22;
                var width = 22;
                var $useIcon = document.createElementNS(svgns, 'use');
                $useIcon.setAttributeNS(xlinkns, 'href', str === '{R}' ? '#RelayIcon' : '#DeprecatedIcon');
                $useIcon.setAttribute('width', width + "px");
                $useIcon.setAttribute('height', height + "px");
                //FIXME: remove hardcoded offset
                var y = parseInt($iconPlaceholder.getAttribute('y')) - 15;
                $useIcon.setAttribute('x', $iconPlaceholder.getAttribute('x'));
                $useIcon.setAttribute('y', y.toString());
                $field.replaceChild($useIcon, $iconPlaceholder);
                continue;
            }
            texts[i].classList.add('field-type');
            if (edgesSources[$field.id] && !/[\[\]\!]/.test(str))
                texts[i].classList.add('type-link');
        }
    });
    forEachNode(svg, '.derived-type', function ($derivedType) {
        $derivedType.classList.add('edge-source');
        $derivedType.querySelector('text').classList.add('type-link');
    });
    forEachNode(svg, '.possible-type', function ($possibleType) {
        $possibleType.classList.add('edge-source');
        $possibleType.querySelector('text').classList.add('type-link');
    });
    var serializer = new XMLSerializer();
    return serializer.serializeToString(svg);
}

// CONCATENATED MODULE: ./src/graph/index.ts





// EXTERNAL MODULE: external "prop-types"
var external_prop_types_ = __webpack_require__(3);

// EXTERNAL MODULE: external "@material-ui/core/styles"
var styles_ = __webpack_require__(7);

// EXTERNAL MODULE: external "@material-ui/core/colors/cyan"
var cyan_ = __webpack_require__(18);
var cyan_default = /*#__PURE__*/__webpack_require__.n(cyan_);

// EXTERNAL MODULE: external "@material-ui/core/colors/yellow"
var yellow_ = __webpack_require__(19);
var yellow_default = /*#__PURE__*/__webpack_require__.n(yellow_);

// EXTERNAL MODULE: ./src/components/variables.css
var variables = __webpack_require__(5);
var variables_default = /*#__PURE__*/__webpack_require__.n(variables);

// CONCATENATED MODULE: ./src/components/MUITheme.tsx
var MUITheme_a;




var theme = Object(styles_["createMuiTheme"])({
    palette: {
        primary: cyan_default.a,
        secondary: yellow_default.a,
    },
    typography: {
        fontSize: 12,
        useNextVariants: true,
    },
    overrides: {
        MuiCheckbox: {
            root: {
                width: '30px',
                height: '15px',
                padding: 0,
            },
        },
        MuiIconButton: {
            root: {
                width: variables_default.a.iconsSize,
                height: variables_default.a.iconSize,
                padding: 0,
            },
        },
        MuiInput: {
            root: {
                marginBottom: '10px',
            },
        },
        MuiTooltip: {
            tooltip: {
                fontSize: variables_default.a.baseFontSize - 2,
            },
        },
        MuiSnackbar: {
            anchorOriginBottomLeft: (MUITheme_a = {},
                MUITheme_a[variables_default.a.bigViewport] = {
                    left: '340px',
                    right: '20px',
                    bottom: '20px',
                },
                MUITheme_a),
        },
        MuiSnackbarContent: {
            root: {
                width: '50%',
                backgroundColor: variables_default.a.alertColor,
            },
        },
    },
});

// EXTERNAL MODULE: external "classnames"
var external_classnames_ = __webpack_require__(2);

// EXTERNAL MODULE: ./src/components/utils/LoadingAnimation.css
var utils_LoadingAnimation = __webpack_require__(34);

// CONCATENATED MODULE: ./src/components/icons/logo-with-signals.svg
var _extends = Object.assign || function (target) {
  for (var i = 1; i < arguments.length; i++) {
    var source = arguments[i];

    for (var key in source) {
      if (Object.prototype.hasOwnProperty.call(source, key)) {
        target[key] = source[key];
      }
    }
  }

  return target;
};

function _objectWithoutProperties(obj, keys) {
  var target = {};

  for (var i in obj) {
    if (keys.indexOf(i) >= 0) continue;
    if (!Object.prototype.hasOwnProperty.call(obj, i)) continue;
    target[i] = obj[i];
  }

  return target;
}


/* harmony default export */ var logo_with_signals = (function (_ref) {
  var _ref$styles = _ref.styles,
      styles = _ref$styles === void 0 ? {} : _ref$styles,
      props = _objectWithoutProperties(_ref, ["styles"]);

  return external_react_default.a.createElement("svg", _extends({
    xmlns: "http://www.w3.org/2000/svg",
    viewBox: "0 0 490.8 438.1"
  }, props), external_react_default.a.createElement("path", {
    d: "M334.2 285c-2.3-2.3-6.1-2.3-8.5 0l-6.5 6.5-10.1-10.1 2.9-2.9c4.7-4.7 4.7-12.3 0-17l-2.6-2.6c.2-.6.4-1.3.4-2.1V217c22.9 15.1 46.9 23.5 67.7 23.5 4.8 0 9.5-.5 13.9-1.4 4.3-.9 7.8-4.1 9.1-8.3 1.3-4.2.1-8.8-3-11.9l-53.1-53.1v-24.3c1.2.2 2.5.4 3.7.4 5.1 0 9.9-2 13.5-5.6 7.5-7.5 7.5-19.6 0-27.1-3.6-3.6-8.4-5.6-13.5-5.6s-9.9 2-13.5 5.6c-3.6 3.6-5.6 8.4-5.6 13.5 0 1.3.1 2.5.4 3.7h-24.3L252 73.3c-2.3-2.3-5.3-3.5-8.5-3.5-1.1 0-2.3.2-3.4.5-4.2 1.3-7.4 4.7-8.3 9.1-4.9 23.5 3.4 53.2 22.2 81.6h-39.8c-.5 0-1 .1-1.5.2 3.3-5.3 2.7-12.4-1.9-17-1.9-1.9-4.2-3.1-6.7-3.7L177.6 30c.4-.3.8-.7 1.2-1.1 5.5-5.5 5.5-14.5 0-20s-14.5-5.5-20 0-5.5 14.5 0 20c2 2 4.5 3.3 7.1 3.8l26.2 110.6c-.3.3-.7.6-1 .9-5.3 5.3-5.3 14 0 19.3.3.3.7.6 1 .9L159.8 197c-.3-.4-.6-.7-.9-1-5.3-5.3-14-5.3-19.3 0-.3.3-.6.7-.9 1L28 170.7c-.6-2.7-1.9-5.1-3.8-7.1-5.5-5.5-14.5-5.5-20 0-2.7 2.7-4.1 6.2-4.1 10s1.5 7.3 4.1 10c2.8 2.8 6.4 4.1 10 4.1 3.6 0 7.2-1.4 10-4.1.4-.4.7-.8 1.1-1.2L136 208.6c.5 2.4 1.8 4.8 3.7 6.7 2.7 2.7 6.2 4 9.7 4 2.1 0 4.2-.5 6.1-1.5l.3.3 99.9 99.9c2.3 2.3 5.3 3.5 8.5 3.5s6.2-1.3 8.5-3.5l3.4-3.4 10.1 10.1-7.4 7.4c-2.3 2.3-2.3 6.1 0 8.5l95.7 95.7c.3.3.6.5.9.8h.1c.3.2.6.4 1 .5h.1c.3.1.7.2 1 .3h.1c.4.1.7.1 1.1.1s.7 0 1.1-.1h.1c.4-.1.7-.2 1-.3h.1c.3-.1.7-.3 1-.5h.1c.7-.4 1.2-1 1.7-1.7v-.1c.2-.3.4-.6.5-1v-.1c.1-.3.2-.7.3-1v-.1c.1-.4.1-.7.1-1.1v-40.9h40.9c1.7 0 3.3-.7 4.4-1.9 0 0 .1 0 .1-.1 2.3-2.3 2.3-6.1 0-8.5l-96-95.6zm38.3 94.2h-34.9v-34.9h34.9v34.9zM325.7 302v30.3h-30.3l30.3-30.3zm0 42.3v26.4l-26.4-26.4h26.4zm12-12v-26.9l26.9 26.9h-26.9zm-69.3-114.4h29.3v29.3l-29.3-29.3zm-3.5-42L295 206h-30.1v-30.1zm63.9-25.5l3.4-3.4v6.8l-3.4-3.4zm14.2-32.7c1.3-1.3 3.1-2.1 5-2.1s3.7.7 5 2.1c2.8 2.8 2.8 7.3 0 10.1-1.3 1.3-3.1 2.1-5 2.1s-3.7-.7-5-2.1-2.1-3.1-2.1-5 .7-3.7 2.1-5.1zm-19.2 20.8l-3.4 3.4-3.4-3.4h6.8zm-80.5-56.6l145.6 145.6c-3.6.7-7.5 1.1-11.5 1.1-21.7 0-48.5-10.8-73.4-30.5l-31.4-31.4c-23.2-29.4-34.1-61.6-29.3-84.8zm9.6 120.5L223.5 173h29.3v29.4zm-88.8 7.2l39.5-39.5 99.9 99.9-5.4 5.4c-.8.3-1.5.7-2.1 1.3-.6.6-1 1.3-1.3 2.1l-21.3 21.3c-.8.3-1.5.7-2.1 1.3-.6.6-1 1.3-1.3 2.1l-6 6-99.9-99.9zm120.3 96.5l16.3-16.3 10.1 10.1-16.3 16.3-10.1-10.1zm88.1 111.5l-26.4-26.4h26.4v26.4zm12-65.3l26.9 26.9h-26.9v-26.9z"
  }), external_react_default.a.createElement("path", {
    d: "M217.5 198.2c-5.2 0-10.2 1.5-14.5 4.5-11.8 8-14.9 24.1-6.9 35.9 4.8 7.1 12.8 11.4 21.4 11.4 5.2 0 10.2-1.5 14.5-4.5 5.7-3.9 9.6-9.7 10.9-16.5 1.3-6.8-.1-13.7-4-19.4-4.8-7.2-12.8-11.4-21.4-11.4zm7.8 37.3c-2.4 1.6-5.1 2.4-7.8 2.4-4.4 0-8.8-2.1-11.5-6.1-4.3-6.3-2.6-14.9 3.7-19.2 2.4-1.6 5.1-2.4 7.8-2.4 4.4 0 8.8 2.1 11.5 6.1 4.2 6.3 2.6 14.9-3.7 19.2z"
  }), external_react_default.a.createElement("path", {
    className: styles["voyager-signal1"] || "voyager-signal1",
    d: "M369.5 101.3c5.1 5.1 10.3 9.4 14.6 12.2 2.6 1.7 5.9 3.6 8.9 3.6 1.4 0 2.6-.4 3.7-1.4 5.8-5.8-8.1-20.9-14.3-27.2-5.1-5.1-10.3-9.4-14.6-12.2-3.8-2.5-9.3-5.4-12.5-2.1-5.9 5.7 8 20.8 14.2 27.1zm7.9-7.9c6.6 6.6 10.6 12 12.4 15.4-3.3-1.9-8.8-5.9-15.4-12.4-6.6-6.6-10.6-12-12.4-15.4 3.4 1.8 8.8 5.8 15.4 12.4z"
  }), external_react_default.a.createElement("path", {
    className: styles["voyager-signal2"] || "voyager-signal2",
    d: "M390 80.8c4.4 4.4 24.2 23.6 34.7 23.6 1.7 0 3.1-.5 4.2-1.6 8.1-8.1-16.9-33.9-22-39-5.1-5.1-30.9-30.1-39-22-3.1 3.1-2 8.4 3.5 16.8 4.4 6.7 11 14.5 18.6 22.2zm12-12.1c13.9 13.9 20.4 24.6 21.4 28.5-3.9-1.1-14.6-7.5-28.5-21.4-13.9-13.9-20.4-24.6-21.4-28.5 4 1.1 14.6 7.5 28.5 21.4z"
  }), external_react_default.a.createElement("path", {
    className: styles["voyager-signal3"] || "voyager-signal3",
    d: "M462.5 67c-6.5-9.2-16.2-20.4-27.3-31.5-11-11-22.2-20.7-31.4-27.2-11.5-8.1-18-10.1-21.5-6.6s-1.5 9.9 6.6 21.5c6.5 9.2 16.2 20.4 27.3 31.5 11.1 11.1 22.3 20.8 31.5 27.3 7.9 5.6 13.4 8.3 17.2 8.3 1.7 0 3.1-.6 4.2-1.7 3.5-3.6 1.6-10-6.6-21.6zm-41.4-17.3C399.4 28 389.3 12.4 387.8 7.2c5.2 1.5 20.8 11.6 42.5 33.3C452 62.2 462 77.8 463.6 83c-5.2-1.6-20.8-11.6-42.5-33.3z"
  }));
});
// CONCATENATED MODULE: ./src/components/utils/LoadingAnimation.tsx
var __extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();




var LoadingAnimation_LoadingAnimation = /** @class */ (function (_super) {
    __extends(LoadingAnimation, _super);
    function LoadingAnimation() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    LoadingAnimation.prototype.render = function () {
        var loading = this.props.loading;
        return (external_react_["createElement"]("div", { className: external_classnames_({ 'loading-box': true, visible: loading }) },
            external_react_["createElement"]("span", { className: "loading-animation" },
                external_react_["createElement"](logo_with_signals, null),
                external_react_["createElement"]("h1", null, " Transmitting... "))));
    };
    return LoadingAnimation;
}(external_react_["Component"]));
/* harmony default export */ var components_utils_LoadingAnimation = (LoadingAnimation_LoadingAnimation);

// CONCATENATED MODULE: ./src/components/GraphViewport.tsx
var GraphViewport_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();



var GraphViewport_GraphViewport = /** @class */ (function (_super) {
    GraphViewport_extends(GraphViewport, _super);
    function GraphViewport() {
        var _this = _super !== null && _super.apply(this, arguments) || this;
        _this.state = { typeGraph: null, displayOptions: null, svgViewport: null };
        // Handle async graph rendering based on this example
        // https://gist.github.com/bvaughn/982ab689a41097237f6e9860db7ca8d6
        _this._currentTypeGraph = null;
        _this._currentDisplayOptions = null;
        return _this;
    }
    GraphViewport.getDerivedStateFromProps = function (props, state) {
        var typeGraph = props.typeGraph, displayOptions = props.displayOptions;
        if (typeGraph !== state.typeGraph || displayOptions !== state.displayOptions) {
            return { typeGraph: typeGraph, displayOptions: displayOptions, svgViewport: null };
        }
        return null;
    };
    GraphViewport.prototype.componentDidMount = function () {
        var _a = this.props, typeGraph = _a.typeGraph, displayOptions = _a.displayOptions;
        this._renderSvgAsync(typeGraph, displayOptions);
    };
    GraphViewport.prototype.componentDidUpdate = function (prevProps, prevState) {
        var svgViewport = this.state.svgViewport;
        if (svgViewport == null) {
            var _a = this.props, typeGraph = _a.typeGraph, displayOptions = _a.displayOptions;
            this._renderSvgAsync(typeGraph, displayOptions);
            return;
        }
        var isJustRendered = prevState.svgViewport == null;
        var _b = this.props, selectedTypeID = _b.selectedTypeID, selectedEdgeID = _b.selectedEdgeID;
        if (prevProps.selectedTypeID !== selectedTypeID || isJustRendered) {
            svgViewport.selectNodeById(selectedTypeID);
        }
        if (prevProps.selectedEdgeID !== selectedEdgeID || isJustRendered) {
            svgViewport.selectEdgeById(selectedEdgeID);
        }
    };
    GraphViewport.prototype.componentWillUnmount = function () {
        this._currentTypeGraph = null;
        this._currentDisplayOptions = null;
        this._cleanupSvgViewport();
    };
    GraphViewport.prototype._renderSvgAsync = function (typeGraph, displayOptions) {
        var _this = this;
        if (typeGraph == null || displayOptions == null) {
            return; // Nothing to render
        }
        if (typeGraph === this._currentTypeGraph && displayOptions === this._currentDisplayOptions) {
            return; // Already rendering in background
        }
        this._currentTypeGraph = typeGraph;
        this._currentDisplayOptions = displayOptions;
        var _a = this.props, svgRenderer = _a.svgRenderer, onSelectNode = _a.onSelectNode, onSelectEdge = _a.onSelectEdge;
        svgRenderer
            .renderSvg(typeGraph, displayOptions)
            .then(function (svg) {
            if (typeGraph !== _this._currentTypeGraph ||
                displayOptions !== _this._currentDisplayOptions) {
                return; // One of the past rendering jobs finished
            }
            _this._cleanupSvgViewport();
            var containerRef = _this.refs['viewport'];
            var svgViewport = new viewport_Viewport(svg, containerRef, onSelectNode, onSelectEdge);
            _this.setState({ svgViewport: svgViewport });
        })
            .catch(function (error) {
            _this._currentTypeGraph = null;
            _this._currentDisplayOptions = null;
            error.message = error.message || 'Unknown error';
            _this.setState(function () {
                throw error;
            });
        });
    };
    GraphViewport.prototype.render = function () {
        var isLoading = this.state.svgViewport == null;
        return (external_react_["createElement"](external_react_["Fragment"], null,
            external_react_["createElement"]("div", { ref: "viewport", className: "viewport" }),
            external_react_["createElement"](components_utils_LoadingAnimation, { loading: isLoading })));
    };
    GraphViewport.prototype.resize = function () {
        var svgViewport = this.state.svgViewport;
        if (svgViewport) {
            svgViewport.resize();
        }
    };
    GraphViewport.prototype.focusNode = function (id) {
        var svgViewport = this.state.svgViewport;
        if (svgViewport) {
            svgViewport.focusElement(id);
        }
    };
    GraphViewport.prototype._cleanupSvgViewport = function () {
        var svgViewport = this.state.svgViewport;
        if (svgViewport) {
            svgViewport.destroy();
        }
    };
    return GraphViewport;
}(external_react_["Component"]));
/* harmony default export */ var components_GraphViewport = (GraphViewport_GraphViewport);

// EXTERNAL MODULE: ./src/components/doc-explorer/TypeList.css
var doc_explorer_TypeList = __webpack_require__(35);

// EXTERNAL MODULE: ./src/components/doc-explorer/TypeLink.css
var doc_explorer_TypeLink = __webpack_require__(36);

// CONCATENATED MODULE: ./src/components/doc-explorer/TypeLink.tsx
var TypeLink_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();





var TypeLink_TypeLink = /** @class */ (function (_super) {
    TypeLink_extends(TypeLink, _super);
    function TypeLink() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    TypeLink.prototype.render = function () {
        var _a = this.props, type = _a.type, onClick = _a.onClick, filter = _a.filter;
        var className;
        if (isBuiltInScalarType(type))
            className = '-built-in';
        else if (isScalarType(type))
            className = '-scalar';
        else if (isInputObjectType(type))
            className = '-input-obj';
        else
            className = '-object';
        return (external_react_["createElement"]("a", { className: external_classnames_('type-name', className), onClick: function (event) {
                event.stopPropagation();
                onClick(type);
            } }, highlightTerm(type.name, filter)));
    };
    return TypeLink;
}(external_react_["Component"]));
/* harmony default export */ var components_doc_explorer_TypeLink = (TypeLink_TypeLink);

// EXTERNAL MODULE: ./src/components/doc-explorer/Description.css
var doc_explorer_Description = __webpack_require__(37);

// EXTERNAL MODULE: external "commonmark"
var external_commonmark_ = __webpack_require__(10);

// CONCATENATED MODULE: ./src/components/utils/Markdown.tsx
var Markdown_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();


var Markdown_Markdown = /** @class */ (function (_super) {
    Markdown_extends(Markdown, _super);
    function Markdown(props) {
        var _this = _super.call(this, props) || this;
        _this.renderer = new external_commonmark_["HtmlRenderer"]({ safe: true });
        _this.parser = new external_commonmark_["Parser"]();
        return _this;
    }
    Markdown.prototype.shouldComponentUpdate = function (nextProps) {
        return this.props.text !== nextProps.text;
    };
    Markdown.prototype.render = function () {
        var _a = this.props, text = _a.text, className = _a.className;
        if (!text)
            return null;
        var parsed = this.parser.parse(text);
        var html = this.renderer.render(parsed);
        return external_react_["createElement"]("div", { className: className, dangerouslySetInnerHTML: { __html: html } });
    };
    return Markdown;
}(external_react_["Component"]));
/* harmony default export */ var utils_Markdown = (Markdown_Markdown);

// CONCATENATED MODULE: ./src/components/doc-explorer/Description.tsx
var Description_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();




var Description_Description = /** @class */ (function (_super) {
    Description_extends(Description, _super);
    function Description() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Description.prototype.render = function () {
        var _a = this.props, text = _a.text, className = _a.className;
        if (text)
            return external_react_["createElement"](utils_Markdown, { text: text, className: external_classnames_('description-box', className) });
        return (external_react_["createElement"]("div", { className: external_classnames_('description-box', className, '-no-description') },
            external_react_["createElement"]("p", null, "No Description")));
    };
    return Description;
}(external_react_["Component"]));
/* harmony default export */ var components_doc_explorer_Description = (Description_Description);

// EXTERNAL MODULE: external "@material-ui/core/IconButton"
var IconButton_ = __webpack_require__(4);
var IconButton_default = /*#__PURE__*/__webpack_require__.n(IconButton_);

// CONCATENATED MODULE: ./src/components/icons/remove-red-eye.svg
var remove_red_eye_extends = Object.assign || function (target) {
  for (var i = 1; i < arguments.length; i++) {
    var source = arguments[i];

    for (var key in source) {
      if (Object.prototype.hasOwnProperty.call(source, key)) {
        target[key] = source[key];
      }
    }
  }

  return target;
};

function remove_red_eye_objectWithoutProperties(obj, keys) {
  var target = {};

  for (var i in obj) {
    if (keys.indexOf(i) >= 0) continue;
    if (!Object.prototype.hasOwnProperty.call(obj, i)) continue;
    target[i] = obj[i];
  }

  return target;
}


/* harmony default export */ var remove_red_eye = (function (_ref) {
  var _ref$styles = _ref.styles,
      styles = _ref$styles === void 0 ? {} : _ref$styles,
      props = remove_red_eye_objectWithoutProperties(_ref, ["styles"]);

  return external_react_default.a.createElement("svg", remove_red_eye_extends({
    height: "24",
    viewBox: "0 0 24 24",
    width: "24",
    xmlns: "http://www.w3.org/2000/svg"
  }, props), external_react_default.a.createElement("path", {
    d: "M0 0h24v24H0z",
    fill: "none"
  }), external_react_default.a.createElement("path", {
    d: "M12 4.5C7 4.5 2.73 7.61 1 12c1.73 4.39 6 7.5 11 7.5s9.27-3.11 11-7.5c-1.73-4.39-6-7.5-11-7.5zM12 17c-2.76 0-5-2.24-5-5s2.24-5 5-5 5 2.24 5 5-2.24 5-5 5zm0-8c-1.66 0-3 1.34-3 3s1.34 3 3 3 3-1.34 3-3-1.34-3-3-3z"
  }));
});
// EXTERNAL MODULE: ./src/components/doc-explorer/FocusTypeButton.css
var FocusTypeButton = __webpack_require__(38);

// CONCATENATED MODULE: ./src/components/doc-explorer/FocusTypeButton.tsx




function FocusTypeButton_FocusTypeButton(props) {
    return (external_react_["createElement"](IconButton_default.a, { className: "eye-button", onClick: props.onClick, color: "primary" },
        external_react_["createElement"](remove_red_eye, null)));
}
/* harmony default export */ var doc_explorer_FocusTypeButton = (FocusTypeButton_FocusTypeButton);

// CONCATENATED MODULE: ./src/components/doc-explorer/TypeList.tsx
var TypeList_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();








var TypeList_TypeList = /** @class */ (function (_super) {
    TypeList_extends(TypeList, _super);
    function TypeList() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    TypeList.prototype.render = function () {
        var _a = this.props, typeGraph = _a.typeGraph, filter = _a.filter, onFocusType = _a.onFocusType, onTypeLink = _a.onTypeLink;
        if (typeGraph === null)
            return null;
        var rootType = typeGraph.nodes[typeGraph.rootId];
        var types = external_lodash_(typeGraph.nodes)
            .values()
            .reject({ id: rootType && rootType.id })
            .sortBy('name')
            .value();
        return (external_react_["createElement"]("div", { className: "doc-explorer-type-list" },
            rootType && renderItem(rootType, '-root'),
            external_lodash_["map"](types, function (type) { return renderItem(type, ''); })));
        function renderItem(type, className) {
            if (!isMatch(type.name, filter)) {
                return null;
            }
            return (external_react_["createElement"]("div", { key: type.id, className: external_classnames_('typelist-item', className) },
                external_react_["createElement"](components_doc_explorer_TypeLink, { type: type, onClick: onTypeLink, filter: filter }),
                external_react_["createElement"](doc_explorer_FocusTypeButton, { onClick: function () { return onFocusType(type); } }),
                external_react_["createElement"](components_doc_explorer_Description, { className: "-doc-type", text: type.description })));
        }
    };
    return TypeList;
}(external_react_["Component"]));
/* harmony default export */ var components_doc_explorer_TypeList = (TypeList_TypeList);

// EXTERNAL MODULE: ./src/components/doc-explorer/TypeDoc.css
var doc_explorer_TypeDoc = __webpack_require__(39);

// EXTERNAL MODULE: external "@material-ui/core/Tooltip"
var Tooltip_ = __webpack_require__(20);
var Tooltip_default = /*#__PURE__*/__webpack_require__.n(Tooltip_);

// EXTERNAL MODULE: ./src/components/doc-explorer/WrappedTypeName.css
var doc_explorer_WrappedTypeName = __webpack_require__(40);

// CONCATENATED MODULE: ./src/components/icons/relay-icon.svg
var relay_icon_extends = Object.assign || function (target) {
  for (var i = 1; i < arguments.length; i++) {
    var source = arguments[i];

    for (var key in source) {
      if (Object.prototype.hasOwnProperty.call(source, key)) {
        target[key] = source[key];
      }
    }
  }

  return target;
};

function relay_icon_objectWithoutProperties(obj, keys) {
  var target = {};

  for (var i in obj) {
    if (keys.indexOf(i) >= 0) continue;
    if (!Object.prototype.hasOwnProperty.call(obj, i)) continue;
    target[i] = obj[i];
  }

  return target;
}


/* harmony default export */ var relay_icon = (function (_ref) {
  var _ref$styles = _ref.styles,
      styles = _ref$styles === void 0 ? {} : _ref$styles,
      props = relay_icon_objectWithoutProperties(_ref, ["styles"]);

  return external_react_default.a.createElement("svg", relay_icon_extends({
    viewBox: "0 0 600 600"
  }, props), external_react_default.a.createElement("g", {
    fill: "#F26B00"
  }, external_react_default.a.createElement("path", {
    d: "M142.536 198.858c0 26.36-21.368 47.72-47.72 47.72-26.36 0-47.722-21.36-47.722-47.72s21.36-47.72 47.72-47.72c26.355 0 47.722 21.36 47.722 47.72"
  }), external_react_default.a.createElement("path", {
    d: "M505.18 414.225H238.124c-35.25 0-63.926-28.674-63.926-63.923s28.678-63.926 63.926-63.926h120.78c20.816 0 37.753-16.938 37.753-37.756s-16.938-37.756-37.753-37.756H94.81c-7.227 0-13.086-5.86-13.086-13.085 0-7.227 5.86-13.086 13.085-13.086h264.093c35.25 0 63.923 28.678 63.923 63.926s-28.674 63.923-63.923 63.923h-120.78c-20.82 0-37.756 16.938-37.756 37.76 0 20.816 16.938 37.753 37.756 37.753H505.18c7.227 0 13.086 5.86 13.086 13.085 0 7.226-5.858 13.085-13.085 13.085z"
  }), external_react_default.a.createElement("path", {
    d: "M457.464 401.142c0-26.36 21.36-47.72 47.72-47.72s47.72 21.36 47.72 47.72-21.36 47.72-47.72 47.72-47.72-21.36-47.72-47.72"
  })));
});
// CONCATENATED MODULE: ./src/components/doc-explorer/WrappedTypeName.tsx
var WrappedTypeName_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();







var WrappedTypeName_WrappedTypeName = /** @class */ (function (_super) {
    WrappedTypeName_extends(WrappedTypeName, _super);
    function WrappedTypeName() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    WrappedTypeName.prototype.render = function () {
        var _a = this.props, container = _a.container, onTypeLink = _a.onTypeLink;
        var type = container.type;
        var wrappers = container.typeWrappers || [];
        var _b = stringifyWrappers(wrappers), leftWrap = _b[0], rightWrap = _b[1];
        return (external_react_["createElement"]("span", { className: "wrapped-type-name" },
            leftWrap,
            external_react_["createElement"](components_doc_explorer_TypeLink, { type: type, onClick: onTypeLink }),
            rightWrap,
            " ",
            container.relayType && wrapRelayIcon()));
    };
    return WrappedTypeName;
}(external_react_["Component"]));
/* harmony default export */ var components_doc_explorer_WrappedTypeName = (WrappedTypeName_WrappedTypeName);
function wrapRelayIcon() {
    return (external_react_["createElement"](Tooltip_default.a, { title: "Relay Connection", placement: "top" },
        external_react_["createElement"](IconButton_default.a, { className: "relay-icon" },
            external_react_["createElement"](relay_icon, null))));
}

// EXTERNAL MODULE: ./src/components/doc-explorer/Argument.css
var doc_explorer_Argument = __webpack_require__(41);

// CONCATENATED MODULE: ./src/components/doc-explorer/Argument.tsx
var Argument_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();





var Argument_Argument = /** @class */ (function (_super) {
    Argument_extends(Argument, _super);
    function Argument() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Argument.prototype.render = function () {
        var _a = this.props, arg = _a.arg, expanded = _a.expanded, onTypeLink = _a.onTypeLink;
        return (external_react_["createElement"]("span", { className: external_classnames_('arg-wrap', { '-expanded': expanded }) },
            external_react_["createElement"]("span", { className: "arg" },
                external_react_["createElement"]("span", { className: "arg-name" }, arg.name),
                external_react_["createElement"](components_doc_explorer_WrappedTypeName, { container: arg, onTypeLink: onTypeLink }),
                arg.defaultValue !== null && (external_react_["createElement"]("span", null,
                    ' = ',
                    external_react_["createElement"]("span", { className: "default-value" }, arg.defaultValue)))),
            external_react_["createElement"](utils_Markdown, { text: arg.description, className: "arg-description" })));
    };
    return Argument;
}(external_react_["Component"]));
/* harmony default export */ var components_doc_explorer_Argument = (Argument_Argument);

// CONCATENATED MODULE: ./src/components/doc-explorer/TypeDoc.tsx
var TypeDoc_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var __assign = (undefined && undefined.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};










var TypeDoc_TypeDoc = /** @class */ (function (_super) {
    TypeDoc_extends(TypeDoc, _super);
    function TypeDoc() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    TypeDoc.prototype.componentDidUpdate = function (prevProps) {
        if (this.props.selectedEdgeID !== prevProps.selectedEdgeID) {
            this.ensureActiveVisible();
        }
    };
    TypeDoc.prototype.componentDidMount = function () {
        this.ensureActiveVisible();
    };
    TypeDoc.prototype.ensureActiveVisible = function () {
        var itemComponent = this.refs['selectedItem'];
        if (!itemComponent)
            return;
        itemComponent.scrollIntoViewIfNeeded();
    };
    TypeDoc.prototype.render = function () {
        var _a = this.props, selectedType = _a.selectedType, selectedEdgeID = _a.selectedEdgeID, typeGraph = _a.typeGraph, filter = _a.filter, onSelectEdge = _a.onSelectEdge, onTypeLink = _a.onTypeLink;
        return (external_react_["createElement"](external_react_["Fragment"], null,
            external_react_["createElement"](components_doc_explorer_Description, { className: "-doc-type", text: selectedType.description }),
            renderTypesDef(selectedType, selectedEdgeID),
            renderFields(selectedType, selectedEdgeID)));
        function renderTypesDef(type, selectedId) {
            var typesTitle;
            var types;
            switch (type.kind) {
                case 'UNION':
                    typesTitle = 'possible types';
                    types = type.possibleTypes;
                    break;
                case 'INTERFACE':
                    typesTitle = 'implementations';
                    types = type.derivedTypes;
                    break;
                case 'OBJECT':
                    typesTitle = 'implements';
                    types = type.interfaces;
                    break;
                default:
                    return null;
            }
            types = types.filter(function (_a) {
                var type = _a.type;
                return typeGraph.nodes[type.id] && isMatch(type.name, filter);
            });
            if (types.length === 0)
                return null;
            return (external_react_["createElement"]("div", { className: "doc-category" },
                external_react_["createElement"]("div", { className: "title" }, typesTitle),
                external_lodash_["map"](types, function (type) {
                    var props = {
                        key: type.id,
                        className: external_classnames_('item', {
                            '-selected': type.id === selectedId,
                        }),
                        onClick: function () { return onSelectEdge(type.id); },
                    };
                    if (type.id === selectedId)
                        props.ref = 'selectedItem';
                    return (external_react_["createElement"]("div", __assign({}, props),
                        external_react_["createElement"](components_doc_explorer_TypeLink, { type: type.type, onClick: onTypeLink, filter: filter }),
                        external_react_["createElement"](components_doc_explorer_Description, { text: type.type.description, className: "-linked-type" })));
                })));
        }
        function renderFields(type, selectedId) {
            var fields = Object.values(type.fields);
            fields = fields.filter(function (field) {
                var args = Object.values(field.args);
                var matchingArgs = args.filter(function (arg) { return isMatch(arg.name, filter); });
                return isMatch(field.name, filter) || matchingArgs.length > 0;
            });
            if (fields.length === 0)
                return null;
            return (external_react_["createElement"]("div", { className: "doc-category" },
                external_react_["createElement"]("div", { className: "title" }, "fields"),
                fields.map(function (field) {
                    var props = {
                        key: field.name,
                        className: external_classnames_('item', {
                            '-selected': field.id === selectedId,
                            '-with-args': !external_lodash_["isEmpty"](field.args),
                        }),
                        onClick: function () { return onSelectEdge(field.id); },
                    };
                    if (field.id === selectedId)
                        props.ref = 'selectedItem';
                    return (external_react_["createElement"]("div", __assign({}, props),
                        external_react_["createElement"]("a", { className: "field-name" }, highlightTerm(field.name, filter)),
                        external_react_["createElement"]("span", { className: external_classnames_('args-wrap', {
                                '-empty': external_lodash_["isEmpty"](field.args),
                            }) }, !external_lodash_["isEmpty"](field.args) && (external_react_["createElement"]("span", { key: "args", className: "args" }, external_lodash_["map"](field.args, function (arg) { return (external_react_["createElement"](components_doc_explorer_Argument, { key: arg.name, arg: arg, expanded: field.id === selectedId, onTypeLink: onTypeLink })); })))),
                        external_react_["createElement"](components_doc_explorer_WrappedTypeName, { container: field, onTypeLink: onTypeLink }),
                        field.isDeprecated && external_react_["createElement"]("span", { className: "doc-alert-text" }, " DEPRECATED"),
                        external_react_["createElement"](utils_Markdown, { text: field.description, className: "description-box -field" })));
                })));
        }
    };
    return TypeDoc;
}(external_react_["Component"]));
/* harmony default export */ var components_doc_explorer_TypeDoc = (TypeDoc_TypeDoc);

// EXTERNAL MODULE: ./src/components/doc-explorer/TypeInfoPopover.css
var TypeInfoPopover = __webpack_require__(42);

// CONCATENATED MODULE: ./src/components/icons/close-black.svg
var close_black_extends = Object.assign || function (target) {
  for (var i = 1; i < arguments.length; i++) {
    var source = arguments[i];

    for (var key in source) {
      if (Object.prototype.hasOwnProperty.call(source, key)) {
        target[key] = source[key];
      }
    }
  }

  return target;
};

function close_black_objectWithoutProperties(obj, keys) {
  var target = {};

  for (var i in obj) {
    if (keys.indexOf(i) >= 0) continue;
    if (!Object.prototype.hasOwnProperty.call(obj, i)) continue;
    target[i] = obj[i];
  }

  return target;
}


/* harmony default export */ var close_black = (function (_ref) {
  var _ref$styles = _ref.styles,
      styles = _ref$styles === void 0 ? {} : _ref$styles,
      props = close_black_objectWithoutProperties(_ref, ["styles"]);

  return external_react_default.a.createElement("svg", close_black_extends({
    height: "24",
    viewBox: "0 0 24 24",
    width: "24",
    xmlns: "http://www.w3.org/2000/svg"
  }, props), external_react_default.a.createElement("path", {
    d: "M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"
  }), external_react_default.a.createElement("path", {
    d: "M0 0h24v24H0z",
    fill: "none"
  }));
});
// CONCATENATED MODULE: ./src/components/doc-explorer/TypeDetails.tsx
var TypeDetails_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();





var TypeDetails_TypeDetails = /** @class */ (function (_super) {
    TypeDetails_extends(TypeDetails, _super);
    function TypeDetails() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    TypeDetails.prototype.renderFields = function (type, onTypeLink) {
        if (external_lodash_["isEmpty"](type.inputFields))
            return null;
        return (external_react_["createElement"]("div", { className: "doc-category" },
            external_react_["createElement"]("div", { className: "title" }, "fields"),
            external_lodash_["map"](type.inputFields, function (field) {
                return (external_react_["createElement"]("div", { key: field.id, className: "item" },
                    external_react_["createElement"]("a", { className: "field-name" }, field.name),
                    external_react_["createElement"](components_doc_explorer_WrappedTypeName, { container: field, onTypeLink: onTypeLink }),
                    external_react_["createElement"](utils_Markdown, { text: field.description, className: "description-box -field" })));
            })));
    };
    TypeDetails.prototype.renderEnumValues = function (type) {
        if (external_lodash_["isEmpty"](type.enumValues))
            return null;
        return (external_react_["createElement"]("div", { className: "doc-category" },
            external_react_["createElement"]("div", { className: "title" }, "values"),
            external_lodash_["map"](type.enumValues, function (value) { return (external_react_["createElement"](TypeDetails_EnumValue, { key: value.name, value: value })); })));
    };
    TypeDetails.prototype.render = function () {
        var _a = this.props, type = _a.type, onTypeLink = _a.onTypeLink;
        return (external_react_["createElement"]("div", { className: "type-details" },
            external_react_["createElement"]("header", null,
                external_react_["createElement"]("h3", null, type.name),
                external_react_["createElement"](components_doc_explorer_Description, { className: "-doc-type", text: type.description })),
            external_react_["createElement"]("div", { className: "doc-categories" },
                this.renderFields(type, onTypeLink),
                this.renderEnumValues(type))));
    };
    return TypeDetails;
}(external_react_["Component"]));
/* harmony default export */ var doc_explorer_TypeDetails = (TypeDetails_TypeDetails);
var TypeDetails_EnumValue = /** @class */ (function (_super) {
    TypeDetails_extends(EnumValue, _super);
    function EnumValue() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    EnumValue.prototype.render = function () {
        var value = this.props.value;
        return (external_react_["createElement"]("div", { className: "item" },
            external_react_["createElement"]("div", { className: "enum-value" }, value.name),
            external_react_["createElement"](utils_Markdown, { className: "description-box -enum-value", text: value.description }),
            value.deprecationReason && (external_react_["createElement"](utils_Markdown, { className: "doc-deprecation", text: value.deprecationReason }))));
    };
    return EnumValue;
}(external_react_["Component"]));

// CONCATENATED MODULE: ./src/components/doc-explorer/TypeInfoPopover.tsx
var TypeInfoPopover_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();






var TypeInfoPopover_ScalarDetails = /** @class */ (function (_super) {
    TypeInfoPopover_extends(ScalarDetails, _super);
    function ScalarDetails(props) {
        var _this = _super.call(this, props) || this;
        _this.state = { localType: null };
        return _this;
    }
    ScalarDetails.prototype.close = function () {
        var _this = this;
        this.props.onChange(null);
        setTimeout(function () {
            _this.setState({ localType: null });
        }, 450);
    };
    ScalarDetails.prototype.render = function () {
        var _this = this;
        var _a = this.props, type = _a.type, onChange = _a.onChange;
        //FIXME: implement animation correctly
        //https://facebook.github.io/react/docs/animation.html
        var localType = this.state.localType;
        if (type && (!localType || type.name !== localType.name)) {
            setTimeout(function () {
                _this.setState({ localType: type });
            });
        }
        return (external_react_["createElement"]("div", { className: external_classnames_('type-info-popover', {
                '-opened': !!type,
            }) },
            external_react_["createElement"](IconButton_default.a, { className: "closeButton", onClick: function () { return _this.close(); } },
                external_react_["createElement"](close_black, null)),
            (type || localType) && external_react_["createElement"](doc_explorer_TypeDetails, { type: type || localType, onTypeLink: onChange })));
    };
    return ScalarDetails;
}(external_react_["Component"]));
/* harmony default export */ var doc_explorer_TypeInfoPopover = (TypeInfoPopover_ScalarDetails);

// CONCATENATED MODULE: ./src/components/doc-explorer/OtherSearchResults.tsx
var OtherSearchResults_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();


var OtherSearchResults_OtherSearchResults = /** @class */ (function (_super) {
    OtherSearchResults_extends(OtherSearchResults, _super);
    function OtherSearchResults() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    OtherSearchResults.prototype.render = function () {
        var _a = this.props, typeGraph = _a.typeGraph, withinType = _a.withinType, searchValue = _a.searchValue, onTypeLink = _a.onTypeLink, onFieldLink = _a.onFieldLink;
        var types = Object.values(typeGraph.nodes).filter(function (type) { return type !== withinType; });
        var matchedTypes = [];
        if (withinType != null) {
            var _loop_1 = function (type) {
                if (isMatch(type.name, searchValue)) {
                    matchedTypes.push(external_react_["createElement"]("div", { className: "item", key: type.name, onClick: function () { return onTypeLink(type); } },
                        external_react_["createElement"]("span", { className: "type-name" }, highlightTerm(type.name, searchValue))));
                }
            };
            for (var _i = 0, types_1 = types; _i < types_1.length; _i++) {
                var type = types_1[_i];
                _loop_1(type);
            }
        }
        var matchedFields = [];
        var _loop_2 = function (type) {
            if (matchedFields.length >= 100) {
                return "break";
            }
            if (type.fields == null) {
                return "continue";
            }
            var fields = Object.values(type.fields);
            var _loop_3 = function (field) {
                var args = Object.values(field.args);
                var matchingArgs = args.filter(function (arg) { return isMatch(arg.name, searchValue); });
                if (!isMatch(field.name, searchValue) && matchingArgs.length === 0) {
                    return "continue";
                }
                matchedFields.push(external_react_["createElement"]("div", { className: "item", key: field.id, onClick: function () { return onFieldLink(field, type); } },
                    external_react_["createElement"]("span", { className: "type-name" }, type.name),
                    external_react_["createElement"]("span", { className: "field-name" }, highlightTerm(field.name, searchValue)),
                    matchingArgs.length > 0 && (external_react_["createElement"]("span", { className: "args args-wrap" }, matchingArgs.map(function (arg) { return (external_react_["createElement"]("span", { key: arg.id, className: "arg-wrap" },
                        external_react_["createElement"]("span", { className: "arg arg-name" }, highlightTerm(arg.name, searchValue)))); })))));
            };
            for (var _i = 0, fields_1 = fields; _i < fields_1.length; _i++) {
                var field = fields_1[_i];
                _loop_3(field);
            }
        };
        for (var _b = 0, types_2 = types; _b < types_2.length; _b++) {
            var type = types_2[_b];
            var state_1 = _loop_2(type);
            if (state_1 === "break")
                break;
        }
        if (matchedTypes.length + matchedFields.length === 0) {
            return (external_react_["createElement"]("div", { className: "other-search-results doc-category" },
                external_react_["createElement"]("div", { className: "title" }, "other results"),
                external_react_["createElement"]("div", { className: "doc-alert-text -search" }, "No results found.")));
        }
        return (external_react_["createElement"]("div", { className: "other-search-results doc-category" },
            external_react_["createElement"]("div", { className: "title" }, "other results"),
            matchedTypes,
            matchedFields));
    };
    return OtherSearchResults;
}(external_react_["Component"]));
/* harmony default export */ var doc_explorer_OtherSearchResults = (OtherSearchResults_OtherSearchResults);

// EXTERNAL MODULE: external "@material-ui/core/Input"
var Input_ = __webpack_require__(21);
var Input_default = /*#__PURE__*/__webpack_require__.n(Input_);

// EXTERNAL MODULE: external "@material-ui/core/InputAdornment"
var InputAdornment_ = __webpack_require__(22);
var InputAdornment_default = /*#__PURE__*/__webpack_require__.n(InputAdornment_);

// EXTERNAL MODULE: ./src/components/utils/SearchBox.css
var utils_SearchBox = __webpack_require__(43);

// CONCATENATED MODULE: ./src/components/utils/SearchBox.tsx
var SearchBox_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();




var SearchBox_SearchBox = /** @class */ (function (_super) {
    SearchBox_extends(SearchBox, _super);
    function SearchBox(props) {
        var _this = _super.call(this, props) || this;
        _this.timeout = null;
        _this.handleChange = function (event) {
            var value = event.target.value;
            _this.setState({ value: value });
            clearTimeout(_this.timeout);
            _this.timeout = setTimeout(function () {
                _this.props.onSearch(value);
            }, 200);
        };
        _this.handleClear = function () {
            _this.setState({ value: '' });
            clearTimeout(_this.timeout);
            _this.props.onSearch('');
        };
        _this.state = { value: props.value || '' };
        return _this;
    }
    SearchBox.prototype.componentWillUnmount = function () {
        clearTimeout(this.timeout);
    };
    SearchBox.prototype.render = function () {
        var value = this.state.value;
        var placeholder = this.props.placeholder;
        return (external_react_["createElement"]("div", { className: "search-box-wrapper" },
            external_react_["createElement"](Input_default.a, { fullWidth: true, placeholder: placeholder, value: value, onChange: this.handleChange, type: "text", className: "search-box", inputProps: { 'aria-label': 'Description' }, endAdornment: value && (external_react_["createElement"](InputAdornment_default.a, { position: "end" },
                    external_react_["createElement"]("span", { className: "search-box-clear", onClick: this.handleClear }, "\u00D7"))) })));
    };
    return SearchBox;
}(external_react_["Component"]));
/* harmony default export */ var components_utils_SearchBox = (SearchBox_SearchBox);

// EXTERNAL MODULE: ./src/components/doc-explorer/DocExplorer.css
var doc_explorer_DocExplorer = __webpack_require__(44);

// CONCATENATED MODULE: ./src/components/doc-explorer/DocExplorer.tsx
var DocExplorer_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var DocExplorer_assign = (undefined && undefined.__assign) || function () {
    DocExplorer_assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return DocExplorer_assign.apply(this, arguments);
};









var initialNav = { title: 'Type List', type: null, searchValue: null };
var DocExplorer_DocExplorer = /** @class */ (function (_super) {
    DocExplorer_extends(DocExplorer, _super);
    function DocExplorer() {
        var _this = _super !== null && _super.apply(this, arguments) || this;
        _this.state = { navStack: [initialNav], typeForInfoPopover: null };
        _this.handleSearch = function (value) {
            var navStack = _this.state.navStack.slice();
            var currentNav = navStack[navStack.length - 1];
            navStack[navStack.length - 1] = DocExplorer_assign({}, currentNav, { searchValue: value });
            _this.setState({ navStack: navStack });
        };
        _this.handleTypeLink = function (type) {
            var _a = _this.props, onFocusNode = _a.onFocusNode, onSelectNode = _a.onSelectNode;
            if (type_graph_isNode(type)) {
                onFocusNode(type.id);
                onSelectNode(type.id);
            }
            else {
                _this.setState({ typeForInfoPopover: type });
            }
        };
        _this.handleFieldLink = function (field, type) {
            var _a = _this.props, onFocusNode = _a.onFocusNode, onSelectNode = _a.onSelectNode, onSelectEdge = _a.onSelectEdge;
            onFocusNode(type.id);
            onSelectNode(type.id);
            // wait for docs panel to rerender with new edges
            setTimeout(function () { return onSelectEdge(field.id); });
        };
        _this.handleNavBackClick = function () {
            var _a = _this.props, onFocusNode = _a.onFocusNode, onSelectNode = _a.onSelectNode;
            var newNavStack = _this.state.navStack.slice(0, -1);
            var newCurrentNode = newNavStack[newNavStack.length - 1];
            _this.setState({ navStack: newNavStack, typeForInfoPopover: null });
            if (newCurrentNode.type == null) {
                return onSelectNode(null);
            }
            onFocusNode(newCurrentNode.type.id);
            onSelectNode(newCurrentNode.type.id);
        };
        return _this;
    }
    DocExplorer.getDerivedStateFromProps = function (props, state) {
        var selectedTypeID = props.selectedTypeID, typeGraph = props.typeGraph;
        var navStack = state.navStack;
        var lastNav = navStack[navStack.length - 1];
        var lastTypeID = lastNav.type ? lastNav.type.id : null;
        if (selectedTypeID !== lastTypeID) {
            if (selectedTypeID == null) {
                return { navStack: [initialNav], typeForInfoPopover: null };
            }
            var type = typeGraph.nodes[selectedTypeID];
            var newNavStack = navStack.concat([{ title: type.name, type: type, searchValue: null }]);
            return { navStack: newNavStack, typeForInfoPopover: null };
        }
        return null;
    };
    DocExplorer.prototype.render = function () {
        var _this = this;
        var typeGraph = this.props.typeGraph;
        if (!typeGraph) {
            return (external_react_["createElement"]("div", { className: "type-doc", key: 0 },
                external_react_["createElement"]("span", { className: "loading" }, " Loading... "),
                ";"));
        }
        var navStack = this.state.navStack;
        var previousNav = navStack[navStack.length - 2];
        var currentNav = navStack[navStack.length - 1];
        var name = currentNav.type ? currentNav.type.name : 'Schema';
        return (external_react_["createElement"]("div", { className: "type-doc", key: navStack.length },
            this.renderNavigation(previousNav, currentNav),
            external_react_["createElement"]("div", { className: "scroll-area" },
                external_react_["createElement"](components_utils_SearchBox, { placeholder: "Search " + name + "...", value: currentNav.searchValue, onSearch: this.handleSearch }),
                this.renderCurrentNav(currentNav),
                currentNav.searchValue && (external_react_["createElement"](doc_explorer_OtherSearchResults, { typeGraph: typeGraph, withinType: currentNav.type, searchValue: currentNav.searchValue, onTypeLink: this.handleTypeLink, onFieldLink: this.handleFieldLink }))),
            currentNav.type && (external_react_["createElement"](doc_explorer_TypeInfoPopover, { type: this.state.typeForInfoPopover, onChange: function (type) { return _this.setState({ typeForInfoPopover: type }); } }))));
    };
    DocExplorer.prototype.renderCurrentNav = function (currentNav) {
        var _a = this.props, typeGraph = _a.typeGraph, selectedEdgeID = _a.selectedEdgeID, onSelectEdge = _a.onSelectEdge, onFocusNode = _a.onFocusNode;
        if (currentNav.type) {
            return (external_react_["createElement"](components_doc_explorer_TypeDoc, { selectedType: currentNav.type, selectedEdgeID: selectedEdgeID, typeGraph: typeGraph, filter: currentNav.searchValue, onTypeLink: this.handleTypeLink, onSelectEdge: onSelectEdge }));
        }
        return (external_react_["createElement"](components_doc_explorer_TypeList, { typeGraph: typeGraph, filter: currentNav.searchValue, onTypeLink: this.handleTypeLink, onFocusType: function (type) { return onFocusNode(type.id); } }));
    };
    DocExplorer.prototype.renderNavigation = function (previousNav, currentNav) {
        var onFocusNode = this.props.onFocusNode;
        if (previousNav) {
            return (external_react_["createElement"]("div", { className: "doc-navigation" },
                external_react_["createElement"]("span", { className: "back", onClick: this.handleNavBackClick }, previousNav.title),
                external_react_["createElement"]("span", { className: "active", title: currentNav.title },
                    currentNav.title,
                    external_react_["createElement"](doc_explorer_FocusTypeButton, { onClick: function () { return onFocusNode(currentNav.type.id); } }))));
        }
        return (external_react_["createElement"]("div", { className: "doc-navigation" },
            external_react_["createElement"]("span", { className: "header" }, currentNav.title)));
    };
    return DocExplorer;
}(external_react_["Component"]));
/* harmony default export */ var components_doc_explorer_DocExplorer = (DocExplorer_DocExplorer);

// EXTERNAL MODULE: ./src/components/utils/PoweredBy.css
var utils_PoweredBy = __webpack_require__(45);

// CONCATENATED MODULE: ./src/components/utils/PoweredBy.tsx
var PoweredBy_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();


var PoweredBy_PoweredBy = /** @class */ (function (_super) {
    PoweredBy_extends(PoweredBy, _super);
    function PoweredBy() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    PoweredBy.prototype.render = function () {
        return (external_react_["createElement"]("div", { className: "powered-by" },
            "\uD83D\uDEF0 Powered by",
            ' ',
            external_react_["createElement"]("a", { href: "https://github.com/APIs-guru/graphql-voyager", target: "_blank" }, "GraphQL Voyager")));
    };
    return PoweredBy;
}(external_react_["Component"]));
/* harmony default export */ var components_utils_PoweredBy = (PoweredBy_PoweredBy);

// EXTERNAL MODULE: external "@material-ui/core/Checkbox"
var Checkbox_ = __webpack_require__(6);
var Checkbox_default = /*#__PURE__*/__webpack_require__.n(Checkbox_);

// EXTERNAL MODULE: external "@material-ui/core/Select"
var Select_ = __webpack_require__(23);
var Select_default = /*#__PURE__*/__webpack_require__.n(Select_);

// EXTERNAL MODULE: external "@material-ui/core/MenuItem"
var MenuItem_ = __webpack_require__(11);
var MenuItem_default = /*#__PURE__*/__webpack_require__.n(MenuItem_);

// EXTERNAL MODULE: ./src/components/settings/RootSelector.css
var settings_RootSelector = __webpack_require__(46);

// CONCATENATED MODULE: ./src/components/settings/RootSelector.tsx
var RootSelector_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();





var RootSelector_RootSelector = /** @class */ (function (_super) {
    RootSelector_extends(RootSelector, _super);
    function RootSelector() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    RootSelector.prototype.render = function () {
        var _a = this.props, schema = _a.schema, onChange = _a.onChange;
        var rootType = this.props.rootType || getDefaultRoot(schema);
        var rootTypeNames = getRootTypeNames(schema);
        var otherTypeNames = Object.keys(schema.types)
            .map(function (id) { return schema.types[id]; })
            .filter(type_graph_isNode)
            .map(function (type) { return type.name; })
            .filter(function (name) { return rootTypeNames.indexOf(name) === -1; })
            .sort();
        return (external_react_["createElement"](Select_default.a, { className: "root-selector", onChange: handleChange, value: rootType },
            rootTypeNames.map(function (name) { return (external_react_["createElement"](MenuItem_default.a, { value: name, key: name },
                external_react_["createElement"]("strong", null, name))); }),
            otherTypeNames.map(function (name) { return (external_react_["createElement"](MenuItem_default.a, { value: name, key: name }, name)); })));
        function handleChange(event) {
            var newRootType = event.target.value;
            if (newRootType !== rootType) {
                onChange(newRootType);
            }
        }
    };
    return RootSelector;
}(external_react_["Component"]));
/* harmony default export */ var components_settings_RootSelector = (RootSelector_RootSelector);
function getRootTypeNames(schema) {
    var queryType = schema.queryType, mutationType = schema.mutationType, subscriptionType = schema.subscriptionType;
    var names = [];
    if (queryType) {
        names.push(queryType.name);
    }
    if (mutationType) {
        names.push(mutationType.name);
    }
    if (subscriptionType) {
        names.push(subscriptionType.name);
    }
    return names;
}

// CONCATENATED MODULE: ./src/components/settings/Settings.tsx
var Settings_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();



var Settings_Settings = /** @class */ (function (_super) {
    Settings_extends(Settings, _super);
    function Settings() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Settings.prototype.render = function () {
        var _a = this.props, schema = _a.schema, options = _a.options, onChange = _a.onChange;
        return (external_react_["createElement"]("div", { className: "menu-content" },
            external_react_["createElement"]("div", { className: "setting-change-root" },
                external_react_["createElement"](components_settings_RootSelector, { schema: schema, rootType: options.rootType, onChange: function (rootType) { return onChange({ rootType: rootType }); } })),
            external_react_["createElement"]("div", { className: "setting-other-options" },
                external_react_["createElement"](Checkbox_default.a, { id: "sort", color: "primary", checked: !!options.sortByAlphabet, onChange: function (event) { return onChange({ sortByAlphabet: event.target.checked }); } }),
                external_react_["createElement"]("label", { htmlFor: "sort" }, "Sort by Alphabet"),
                external_react_["createElement"](Checkbox_default.a, { id: "skip", color: "primary", checked: !!options.skipRelay, onChange: function (event) { return onChange({ skipRelay: event.target.checked }); } }),
                external_react_["createElement"]("label", { htmlFor: "skip" }, "Skip Relay"),
                external_react_["createElement"](Checkbox_default.a, { id: "deprecated", color: "primary", checked: !!options.skipDeprecated, onChange: function (event) { return onChange({ skipDeprecated: event.target.checked }); } }),
                external_react_["createElement"]("label", { htmlFor: "deprecated" }, "Skip deprecated"),
                external_react_["createElement"](Checkbox_default.a, { id: "showLeafFields", color: "primary", checked: !!options.showLeafFields, onChange: function (event) { return onChange({ showLeafFields: event.target.checked }); } }),
                external_react_["createElement"]("label", { htmlFor: "showLeafFields" }, "Show leaf fields"))));
    };
    return Settings;
}(external_react_["Component"]));
/* harmony default export */ var settings_Settings = (Settings_Settings);

// EXTERNAL MODULE: ./src/components/Voyager.css
var components_Voyager = __webpack_require__(47);

// EXTERNAL MODULE: ./src/components/viewport.css
var viewport = __webpack_require__(48);

// CONCATENATED MODULE: ./src/components/Voyager.tsx
var Voyager_extends = (undefined && undefined.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var Voyager_assign = (undefined && undefined.__assign) || function () {
    Voyager_assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return Voyager_assign.apply(this, arguments);
};













var defaultDisplayOptions = {
    rootType: undefined,
    skipRelay: true,
    skipDeprecated: true,
    sortByAlphabet: false,
    showLeafFields: true,
    hideRoot: false,
};
function normalizeDisplayOptions(options) {
    return options != null ? Voyager_assign({}, defaultDisplayOptions, options) : defaultDisplayOptions;
}
var Voyager_Voyager = /** @class */ (function (_super) {
    Voyager_extends(Voyager, _super);
    function Voyager(props) {
        var _this = _super.call(this, props) || this;
        _this.state = {
            introspectionData: null,
            schema: null,
            typeGraph: null,
            displayOptions: defaultDisplayOptions,
            selectedTypeID: null,
            selectedEdgeID: null,
        };
        _this.viewportRef = external_react_["createRef"]();
        _this.instospectionPromise = null;
        _this.handleDisplayOptionsChange = function (delta) {
            var displayOptions = Voyager_assign({}, _this.state.displayOptions, delta);
            _this.updateIntrospection(_this.state.introspectionData, displayOptions);
        };
        _this.handleSelectNode = function (selectedTypeID) {
            if (selectedTypeID !== _this.state.selectedTypeID) {
                _this.setState({ selectedTypeID: selectedTypeID, selectedEdgeID: null });
            }
        };
        _this.handleSelectEdge = function (selectedEdgeID) {
            if (selectedEdgeID === _this.state.selectedEdgeID) {
                // deselect if click again
                _this.setState({ selectedEdgeID: null });
            }
            else {
                var selectedTypeID = extractTypeId(selectedEdgeID);
                _this.setState({ selectedTypeID: selectedTypeID, selectedEdgeID: selectedEdgeID });
            }
        };
        _this.svgRenderer = new svg_renderer_SVGRender(_this.props.workerURI, _this.props.loadWorker);
        return _this;
    }
    Voyager.prototype.componentDidMount = function () {
        this.fetchIntrospection();
    };
    Voyager.prototype.fetchIntrospection = function () {
        var _this = this;
        var displayOptions = normalizeDisplayOptions(this.props.displayOptions);
        if (typeof this.props.introspection !== 'function') {
            this.updateIntrospection(this.props.introspection, displayOptions);
            return;
        }
        var promise = this.props.introspection(Object(utilities_["getIntrospectionQuery"])());
        if (!isPromise(promise)) {
            throw new Error('SchemaProvider did not return a Promise for introspection.');
        }
        this.setState({
            introspectionData: null,
            schema: null,
            typeGraph: null,
            displayOptions: null,
            selectedTypeID: null,
            selectedEdgeID: null,
        });
        this.instospectionPromise = promise;
        promise.then(function (introspectionData) {
            if (promise === _this.instospectionPromise) {
                _this.instospectionPromise = null;
                _this.updateIntrospection(introspectionData, displayOptions);
            }
        });
    };
    Voyager.prototype.updateIntrospection = function (introspectionData, displayOptions) {
        var schema = getSchema(introspectionData, displayOptions.sortByAlphabet, displayOptions.skipRelay, displayOptions.skipDeprecated);
        var typeGraph = getTypeGraph(schema, displayOptions.rootType, displayOptions.hideRoot);
        this.setState({
            introspectionData: introspectionData,
            schema: schema,
            typeGraph: typeGraph,
            displayOptions: displayOptions,
            selectedTypeID: null,
            selectedEdgeID: null,
        });
    };
    Voyager.prototype.componentDidUpdate = function (prevProps) {
        if (this.props.introspection !== prevProps.introspection) {
            this.fetchIntrospection();
        }
        else if (this.props.displayOptions !== prevProps.displayOptions) {
            this.updateIntrospection(this.state.introspectionData, normalizeDisplayOptions(this.props.displayOptions));
        }
        if (this.props.hideDocs !== prevProps.hideDocs) {
            this.viewportRef.current.resize();
        }
    };
    Voyager.prototype.render = function () {
        var _a = this.props, _b = _a.hideDocs, hideDocs = _b === void 0 ? false : _b, _c = _a.hideSettings, hideSettings = _c === void 0 ? false : _c;
        return (external_react_["createElement"](styles_["MuiThemeProvider"], { theme: theme },
            external_react_["createElement"]("div", { className: "graphql-voyager" },
                !hideDocs && this.renderPanel(),
                !hideSettings && this.renderSettings(),
                this.renderGraphViewport())));
    };
    Voyager.prototype.renderPanel = function () {
        var _this = this;
        var children = external_react_["Children"].toArray(this.props.children);
        var panelHeader = children.find(function (child) { return child.type === Voyager.PanelHeader; });
        var _a = this.state, typeGraph = _a.typeGraph, selectedTypeID = _a.selectedTypeID, selectedEdgeID = _a.selectedEdgeID;
        var onFocusNode = function (id) { return _this.viewportRef.current.focusNode(id); };
        return (external_react_["createElement"]("div", { className: "doc-panel" },
            external_react_["createElement"]("div", { className: "contents" },
                panelHeader,
                external_react_["createElement"](components_doc_explorer_DocExplorer, { typeGraph: typeGraph, selectedTypeID: selectedTypeID, selectedEdgeID: selectedEdgeID, onFocusNode: onFocusNode, onSelectNode: this.handleSelectNode, onSelectEdge: this.handleSelectEdge }),
                external_react_["createElement"](components_utils_PoweredBy, null))));
    };
    Voyager.prototype.renderSettings = function () {
        var _a = this.state, schema = _a.schema, displayOptions = _a.displayOptions;
        if (schema == null)
            return null;
        return (external_react_["createElement"](settings_Settings, { schema: schema, options: displayOptions, onChange: this.handleDisplayOptionsChange }));
    };
    Voyager.prototype.renderGraphViewport = function () {
        var _a = this.state, displayOptions = _a.displayOptions, typeGraph = _a.typeGraph, selectedTypeID = _a.selectedTypeID, selectedEdgeID = _a.selectedEdgeID;
        return (external_react_["createElement"](components_GraphViewport, { svgRenderer: this.svgRenderer, typeGraph: typeGraph, displayOptions: displayOptions, selectedTypeID: selectedTypeID, selectedEdgeID: selectedEdgeID, onSelectNode: this.handleSelectNode, onSelectEdge: this.handleSelectEdge, ref: this.viewportRef }));
    };
    Voyager.propTypes = {
        introspection: external_prop_types_["oneOfType"]([external_prop_types_["func"].isRequired, external_prop_types_["object"].isRequired])
            .isRequired,
        displayOptions: external_prop_types_["shape"]({
            rootType: external_prop_types_["string"],
            skipRelay: external_prop_types_["bool"],
            skipDeprecated: external_prop_types_["bool"],
            sortByAlphabet: external_prop_types_["bool"],
            hideRoot: external_prop_types_["bool"],
            showLeafFields: external_prop_types_["bool"],
        }),
        hideDocs: external_prop_types_["bool"],
        hideSettings: external_prop_types_["bool"],
        workerURI: external_prop_types_["string"],
        loadWorker: external_prop_types_["func"],
    };
    Voyager.PanelHeader = function (props) {
        return props.children || null;
    };
    return Voyager;
}(external_react_["Component"]));
/* harmony default export */ var src_components_Voyager = (Voyager_Voyager);
// Duck-type promise detection.
function isPromise(value) {
    return typeof value === 'object' && typeof value.then === 'function';
}

// CONCATENATED MODULE: ./src/components/index.tsx


// CONCATENATED MODULE: ./src/index.tsx
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "init", function() { return init; });
/* concated harmony reexport GraphQLVoyager */__webpack_require__.d(__webpack_exports__, "GraphQLVoyager", function() { return src_components_Voyager; });
/* concated harmony reexport Voyager */__webpack_require__.d(__webpack_exports__, "Voyager", function() { return src_components_Voyager; });
var src_assign = (undefined && undefined.__assign) || function () {
    src_assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return src_assign.apply(this, arguments);
};



function init(element, options) {
    external_react_dom_["render"](external_react_["createElement"](src_components_Voyager, src_assign({}, options)), element);
}



/***/ })
/******/ ]);
});