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

	//storge loaded module object
	var moduleMap = {};

	//
	var loadingMap = {};

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
		}
	};

	function addScript(url, attr, loaded) {

		var s = d.createElement("script");
		s.type = "text/javascript";

		for(var k in attr) {
			s.setAttribute(k, attr[k]);
		}
		
		s.onload = loaded;
		s.src = url;
		d.body.appendChild(s);
	}



	function ModuleLoader(modName) {

		var self = this;

		self.name = modName;
		self.alias = '';
		self.factory = null;

		self.onload = function () {};

	}

	ModuleLoader.prototype = {

		load: function () {
			
			var self = this;
			var name = self.name;
			var url = Utils.getJSIntactURL(name);
			addScript(url, {}, function () {
				self.loaded.call(self, this);
			});

		},

		loaded: function (target) {

			var self = this;
			var cache  = Fn.getModuleCache();
			self.factory = cache.factory;
			self.deps = cache.deps;

			Fn.setCacheMap(self.name, self);

			self.loadDeps(self.deps);

		},

		loadDeps: function (deps) {

			var self = this;
			var url = '';
			var loader;
			var num = 0;
			var loadedNum = 0;
			for(var alias in deps) {
				loader = new ModuleLoader(deps[alias]);
				loader.load();
				loader.onload = function () {
					loadedNum++;
					if(num === loadedNum) {
						self.onload()
					}
				}
				num++;
			}

		},

		depsLoaded: function () {

		}

	};

	

	w.define = function(deps, factory) {
		
		if(typeof deps === 'function') {
			factory = deps;
			deps = {};
		}

		Fn.setModuleCache({
			deps: deps,
			factory: factory
		});

	};

	w.require = function(deps, factory) {

		

	};

	//init
	function init() {

		var self = Utils.getSelfElem(),
			selfUrl = self.src,
			mainModName = self.getAttribute('data-main');

		self = null;

		config.basePath = Utils.getBasePath(selfUrl);

		var loader = new ModuleLoader(mainModName);
		loader.onload = function () {
			console.log(cacheMap);
		};
		loader.load();
		
	}

	init();

	//debug
	w.moduleMap = moduleMap;
	w.cacheMap = cacheMap;
	w.loadingMap = loadingMap;

})(window);