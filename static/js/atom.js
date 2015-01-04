/*

module loader

*/

(function(w){

	var d = w.document,
		l = w.location;

	var config = {
		basePath : '',
		jsFileTail : ".js",
		cssFileTail : ".css"
	};

	//
	var runTime = {};

	//storge loaded module object
	var moduleMap = {};

	//
	var moduleCache = null;

	//
	var Utils = {
		getJSIntactURL: function (modName) {
			return config.basePath + module + config.jsFileTail;
		},
		getSelfElem: function () {
			var s = d.getElementsByTagName('script');
			return s[s.length-1];
		},
		getBasePath: function (s) {
			return s.replace(/[^\/]\.js/, '');
		}
	};

	function getModuleCache() {
		var c = moduleCache;
		moduleCache = null;
		return c;
	}

	function Module(modMap, getCacheFunc) {

		var self = this;
		self.modMap = modMap;
		self.getCacheFunc = getCacheFunc;
		self.modAttrNameKey = "data-name";

		self.onload = function () {};
	}

	Module.prototype = {

		pack: function (e) {

			var self = this;
			var modName = self.getAttribute(self.modAttrNameKey);
			self.onload(modName, self.getCacheFunc());
		},

		load: function (url, modName) {

			var mod = self.modMap[modName];
			if(mod) {
				self.onload(modName, mod);
				return;
			}

			var self = this;

			var s = d.createElement("script");
			s.type = "text/javascript";
			s.setAttribute(self.modAttrNameKey, modName);
			s.onload = self.pack;
			s.url = url;
		},

		loadDeps: function () {

		}



	};

	/*function loadModule(module, callback) {

		if(moduleMap[module]) {
			callback(module, moduleMap[module]);
			return;
		}

		var url = Utils.getJSIntactURL(module),
			s = d.createElement("script");

		s.type = "text/javascript";
		s.setAttribute("data-name", module);
		s.src = url;

		s.onload = function(e) {
			
			var t = e.target,
				modName = t.getAttribute('data-name');
			
			moduleMap[modName] = moduleCache;
			callback(modName, moduleCache);
			moduleCache = null;
		};
		d.body.appendChild(s);
	}

	function loadModules(modMap, callback) {
		
		var modNum = 0,
			loadedNum = 0,
			m = {};

		function exeContext(alias, module) {

			loadModule(module, function(k, v) {
				loadedNum++;
				m[alias] = v;
				if(loadedNum === modNum) callback(m);
			});
		}

		for(var i in modMap) {
			modNum++;
			exeContext(i, modMap[i]);
		}
	}
	
	function executeModule(deps, factory) {

		if(typeof deps === 'function') {
			factory = deps;
			moduleCache = factory(runTime, {});
		} else {
			loadModules(deps, function(m) {
				moduleCache = factory(runTime, m);
			});
		}
	}

	w.require = function(deps, factory) {
		executeModule(deps, factory);
	};*/

	w.define = function(deps, factory) {
		
		var mod = new Module(moduleMap, getModuleCache);

		if(typeof deps === 'function') {
			factory = deps;
			moduleCache = factory(runTime, {});
		} else {
			mod.loadDeps(deps, function(m) {
				moduleCache = factory(runTime, m);
			});
		}

	};

	/*w.config = function(conf) {
		for (var i in conf) {
			runTime[i] = conf[i];	
		}
	};*/

	//init
	function init() {

		var self = Utils.getSelfElem(),
			selfUrl = self.src,
			mainMod = self.getAttribute('data-main');

		self = null;

		config.basePath = Utils.getBasePath(selfUrl);

		loadModule(mainMod, function(){});
	}

	init();

})(window);