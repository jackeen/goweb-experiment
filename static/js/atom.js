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
			moduleCache = null;
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

		//console.log(url);
		/*s.onload = function () {
			loaded(s);
		};*/

		s.onreadystatechange = function () {
			var r = s.readyState;
			if(r === 'loaded' || r === 'complete') {
				isLoading = false;
				loaded(s);
				var loo = loadingLoop.shift();
				if(loo) addScript(loo.url, loo.loaded);
			}
		};

		s.src = url;
		d.body.appendChild(s);
	}

	function ModuleLoader(modName, modAlias) {

		var self = this;

		self.name = modName;
		self.alias = modAlias;
		self.factory = function () {};

		self.allDeps = [];
		self.allAlias = [];

		self.depsNum = 0;
		self.loadedDepsNum = 0;

		self.onload = function () {};

	}

	ModuleLoader.prototype = {

		load: function () {
			
			var self = this;
			var name = self.name;
			var url = Utils.getJSIntactURL(name);

			addScript(url, function (target) {
				self.selfReady(target);
			});

		},

		selfReady: function (target) {

			var self = this;
			var cache  = Fn.getModuleCache();
			self.factory = cache.factory;
			self.allDeps = cache.allDeps;
			self.allAlias = cache.allAlias;

			Fn.setCacheMap(self.name, self);

			self.loadDeps();
		},

		loadDeps: function () {

			var self = this;
			var depsList = self.allDeps;
			var aliasList = self.allAlias;
			var len = depsList.length;
			var loader, alias, modules = {};

			if(len === 0) {
				self.onload(self.factory(runTime, {}));
				return;
			}

			for(var i = 0; i < len; i++) {

				alias = aliasList[i];
				moduleName = depsList[i];
				loader = new ModuleLoader(moduleName, alias);
				loader.onload = function (module) {

					self.loadedDepsNum++;
					modules[this.alias] = module;
					if(self.depsNum === self.loadedDepsNum) {
						self.onload(self.factory(runTime, modules));
					}

				};
				loader.load();
				self.depsNum++;

			}
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
	function init() {

		var self = Utils.getSelfElem(),
			selfUrl = self.src,
			mainModName = self.getAttribute('data-main');

		self = null;

		config.basePath = Utils.getBasePath(selfUrl);
		
		w.onload = function () {
			var loader = new ModuleLoader(mainModName, 'main');
			loader.load();
		};
		
	}

	init();

	//debug
	w.moduleMap = moduleMap;
	w.cacheMap = cacheMap;

})(window);