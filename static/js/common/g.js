/*

module loader

*/

(function(w){

	var d = w.document,
		l = w.location;

	var config = {
		origin: '',
		baseDir : '',
		selfModDir: '',
		jsFileTail : ".js",
		cssFileTail : ".css"
	};

	//
	var runTime = {};

	//storge loaded module object
	var moduleMap = {};

	/*	this is a cache for loaded module executed.
		the module of script run complete and dispatch load event to
		loaded callback function, it read the cache value regist module.		
	*/
	var moduleCache = null;

	function getJSIntactURL(module) {
		return config.baseDir + '/' + module + config.jsFileTail;
	}

	function loadModule(module, callback) {

		if(moduleMap[module]) {
			callback(module, moduleMap[module]);
			return;
		}

		var url = getJSIntactURL(module),
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

	
	/**/
	function getSelfElem() {
		var s = d.getElementsByTagName('script');
		return s[s.length-1];
	}

	function exeuteModule(deps, factory) {

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
		exeuteModule(deps, factory);
	};

	w.define = function(deps, factory) {
		exeuteModule(deps, factory);
	};

	w.config = function(conf) {
		for (var i in conf) {
			runTime[i] = conf[i];	
		}
	};

	//init
	function init() {

		var self = getSelfElem(),
			selfUrl = self.src,
			mainMod = self.getAttribute('data-main');

		self = null;
		
		var origin = l.origin,
			path = selfUrl.replace(origin + '/', ''),
			pathArr = path.split('/'),
			jsRoot = pathArr.splice(0, 1)[0],
			gName = pathArr.splice(pathArr.length-1, 1)[0],
			gDir = pathArr.join('/');

		config.origin = origin;
		config.baseDir = '/' + jsRoot;
		config.selfModDir = '/' + gDir;

		loadModule(mainMod, function(){});
		
	}

	init();

})(window);