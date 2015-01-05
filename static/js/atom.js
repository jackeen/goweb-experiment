/*

module loader

*/

(function(w){

	var d = w.document,
		l = w.location;

	var config = {
		basePath : '',
		jsFileTail : ".js",
		cssFileTail : ".css",
		modAttrNameKey: "data-name"
	};

	//
	var runTime = {};

	//storge loaded module object
	var moduleMap = {};

	//
	var loadingMap = {};

	//
	var moduleCache = null;

	//
	var Utils = {
		getJSIntactURL: function (modName) {
			return config.basePath + modName + config.jsFileTail;
		},
		getSelfElem: function () {
			var s = d.getElementsByTagName('script');
			return s[s.length-1];
		},
		getBasePath: function (s) {
			return s.replace(/[^\/]+\.js/, '');
		}
	};

	//get current run factory and clean it
	function getModuleCache() {
		var c = moduleCache;
		moduleCache = null;
		return c;
	}

	function loadModule(name, onload) {

		var loadedMod = moduleMap[name];
		if(loadedMod) {
			onload(name, loadedMod);
			return;
		}

		var loadingMod = loadingMap[name];
		if(loadingMod) {
			loadingMod.push(onload);
		} else {
			loadingMap[name] = [onload];
		}

		var s = d.createElement("script");
		s.type = "text/javascript";
		s.setAttribute(config.modAttrNameKey, name);
		s.onload = function () {
			var mName = this.getAttribute(config.modAttrNameKey);
			var cache = getModuleCache();
			var queue = loadingMap[mName];
			var cb = queue.shift();
			for (;typeof cb === 'function'; cb = queue.shift()) {
				moduleMap[mName] = cache;
				cb(mName, cache);
			}
			this.onload = null;
		};
		s.src = Utils.getJSIntactURL(name);
		d.body.appendChild(s);
	}

	function DepsLoader(deps) {
		
		var self = this;
		self.depsMap = {};
		self.num = 0;
		self.loadedNum = 0;

		self.onload = function () {}

		for(var alias in deps) {
			self.num++;
			loadModule(deps[alias], function (modName, mod) {
				self.depsMap[alias] = mod;
				self.loadedNum++;
				if(self.num === self.loadedNum) self.onload(self.depsMap);
				console.log(self.depsMap);
			});
		}
	}

	DepsLoader.prototype = {

		loaded: function () {

		}

	};


	function loadDeps(deps, onload) {

		var loader = new DepsLoader(deps);
		loader.onload = onload;
	}

	w.define = function(deps, factory) {
		
		if(typeof deps === 'function') {
			factory = deps;
			moduleCache = factory(runTime, {});
		} else {
			loadDeps(deps, function(m) {
				moduleCache = factory(runTime, m);
			});
		}

	};

	//init
	function init() {

		var self = Utils.getSelfElem(),
			selfUrl = self.src,
			mainModName = self.getAttribute('data-main');

		self = null;

		config.basePath = Utils.getBasePath(selfUrl);

		loadModule(mainModName);
	}

	init();

	//debug
	w.moduleMap = moduleMap;
	w.loadingMap = loadingMap;

})(window);