/*

project: atomjs

*/

(function(w){

	var d = w.document,
		l = w.location;

	var config = {
		basePath : "",
		jsFileTail : ".js",
		cssFileTail : ".css",
		modAttrNameKey : "data-name",
		modAliasNameKey : "data-alias"
	};

	//
	var runTime = {};

	//storage loaded module object
	var moduleMap = {};

	//
	var moduleCache = null;

	//
	var cacheMap = {};

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

	var Fn = {
		getModuleCache: function () {
			var c = moduleCache;
			//moduleCache = null;
			return c;
		},
		setModuleCache: function (v) {
			moduleCache = v;
		},
		getModuleMap: function (k) {
			return moduleMap[k];
		},
		setModuleMap: function (k, v) {
			moduleMap[k] = v;
		},
		getCacheMap: function (k) {
			return cacheMap[k];
		},
		setCacheMap: function (k, v) {
			cacheMap[k] = v;
		},
		delCacheMap: function (k) {
			delete cacheMap[k];
		}
	};


	//
	var isLoading = false;
	var loadingLoop = [];

	function addScript(url, loaded) {

		if(isLoading) {
			loadingLoop.push({
				url: url,
				loaded: loaded
			});
			return;
		}

		isLoading = true;

		var s = d.createElement("script");
		s.type = "text/javascript";

		console.log(url);
		//s.onload = loaded;

		s.onreadystatechange = function () {
			var r = s.readyState;
			if(r === 'loaded' || r === 'complete') {
				isLoading = false;
				loaded(s);
				var loo = loadingLoop.shift();
				if(loo) addScript(loo.url, loo.loaded);
			}
		}

		s.src = url;
		d.body.appendChild(s);
	}

	function ModuleLoader(modName) {

		var self = this;

		self.name = modName;
		self.factory = null;
		self.depsNum = 0;
		self.loadedDepsNum = 0;

		self.onload = function () {};

	}

	ModuleLoader.prototype = {

		load: function () {
			
			var self = this;
			var name = self.name;
			var url = Utils.getJSIntactURL(name);

			addScript(url, function () {
				self.loaded.call(self, this);
			});

		},

		loaded: function (target) {

			var self = this;
			var cache  = Fn.getModuleCache();
			self.factory = cache.factory;
			self.allDeps = cache.allDeps;
			self.allAlias = cache.allAlias;

			Fn.setCacheMap(self.name, self);

			if(cache.allAlias.length > 0) {
				self.loadDeps(self.allDeps);
			} else {
				self.depsLoaded({});
			}
		},

		loadDeps: function (deps) {

			var self = this;
			var loader;
			var len = deps.length;

			for(var i = 0; i < len; i++) {

				loader = new ModuleLoader(deps[i]);
				loader.load();

				loader.onload = function (m) {

					self.loadedDepsNum++;
					if(self.depsNum === self.loadedDepsNum) {
						self.depsLoaded(m);
					}

				}

				self.depsNum++;

			}
		},

		depsLoaded: function (m) {
			
			var self = this,
				deps = self.allDeps,
				alias = self.allAlias,
				len = deps.length,
				modules = {},
				moduleName,
				module;

			for(var i = 0; i < len; i++) {

				moduleName = deps[i];
				module = Fn.getCacheMap(moduleName).factory(runTime, m);
				Fn.setModuleMap(moduleName, module);
				Fn.delCacheMap(moduleName);
				modules[alias[i]] = module;
				
				//console.log(moduleName, module);
			}
			
			self.onload(modules);
		}

	};

	w.define = function(deps, factory) {
		
		var allAlias = [];
		var allDeps = [];

		if(typeof deps === 'function') {
			factory = deps;
			deps = {};
		}

		for(var alias in deps) {
			allAlias.push(alias);
			allDeps.push(deps[alias]);
		}

		Fn.setModuleCache({
			allAlias: allAlias,
			allDeps: allDeps,
			factory: factory
		});

	};

	//init

	function loadMainModule(modName) {

		var loader = new ModuleLoader(modName);
		loader.onload = function (modules) {
			var moduleName = this.name;
			Fn.getCacheMap(moduleName).factory(runTime, modules);
			Fn.delCacheMap(moduleName);
		};
		loader.load();
	}

	function init() {

		var self = Utils.getSelfElem(),
			selfUrl = self.src,
			mainModName = self.getAttribute('data-main');

		self = null;

		config.basePath = Utils.getBasePath(selfUrl);
		
		w.onload = function () {
			loadMainModule(mainModName);
		};
		
	}

	init();

	//debug
	w.moduleMap = moduleMap;
	w.cacheMap = cacheMap;

})(window);