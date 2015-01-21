/*

project: atomjs
version: 0.1

*/

(function(w){

	var d = w.document,
		l = w.location;

	var config = {
		basePath : "",
		jsFileTail : ".js",
		cssFileTail : ".css"
	};

	//
	var runTime = {};

	//storage loaded module object
	var moduleMap = {};

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
		},
		getBrowser: function () {
			var u = navigator.userAgent;
			var b = {}, version, name, mb;
			var ie = /(MSIE) ([\d.]+)/;
			var other = /(Chrome|Safari|Opera|Firefox)\/([\d.]+)/;
			mb = u.match(other) || u.match(ie);
			name = mb[1];
			version = parseInt(mb[2], 10);
			b[name] = version;
			b['name'] = name;
			b['version'] = version;
			return b;
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
		}
	};


	//
	var ScriptLoader = {

		isLoading: false,
		loadingLoop: [],

		addScript: function (url, loaded) {

			var s = d.createElement("script");
			s.type = "text/javascript";

			s.onload = loaded;

			s.src = url;
			d.body.appendChild(s);
		},

		addScript4IE: function (url, loaded) {

			var self = this;

			if(self.isLoading) {
				self.loadingLoop.push({
					url: url,
					loaded: loaded
				});
				return;
			}

			self.isLoading = true;

			var s = d.createElement("script");
			s.type = "text/javascript";

			s.onreadystatechange = function () {
				var r = s.readyState;
				if(r === 'loaded' || r === 'complete') {
					self.isLoading = false;
					loaded.call(this);
					var loo = self.loadingLoop.shift();
					if(loo) self.addScript(loo.url, loo.loaded);
				}
			};

			s.src = url;
			d.body.appendChild(s);
		}

	};

	//
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

			ScriptLoader.addScript(url, function () {
				self.selfReady();
			});

		},

		selfReady: function () {

			var self = this;
			var cache  = Fn.getModuleCache();
			self.factory = cache.factory;
			self.allDeps = cache.allDeps;
			self.allAlias = cache.allAlias;

			self.loadDeps();
		},

		loadDeps: function () {

			var self = this;
			var depsList = self.allDeps;
			var aliasList = self.allAlias;
			var len = depsList.length;
			var loader, alias, modules = {}, loadedModule;

			if(len === 0) {
				self.onload(self.factory(runTime, {}));
				return;
			}

			for(var i = 0; i < len; i++) {

				self.depsNum++;
				alias = aliasList[i];
				moduleName = depsList[i];

				loadedModule = Fn.setModuleMap(moduleName);
				if(loadedModule) {
					modules[alias] = loadedModule;
					continue;
				}

				loader = new ModuleLoader(moduleName, alias);
				loader.onload = function (module) {

					self.loadedDepsNum++;
					modules[this.alias] = module;
					if(self.depsNum === self.loadedDepsNum) {
						self.onload(self.factory(runTime, modules));
					}

					Fn.setModuleMap(this.name, module);

				};
				loader.load();

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

		var b = Utils.getBrowser();
		if(b.name === 'MSIE' && b.version <= 9) {
			ScriptLoader.addScript = ScriptLoader.addScript4IE;
		}
		runTime['browser'] = b;
		
		w.onload = function () {
			var loader = new ModuleLoader(mainModName, 'main');
			loader.load();
		};
		
	}

	init();

	//debug
	w.atomModuleMap = moduleMap;

})(window);