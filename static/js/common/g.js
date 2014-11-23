//common

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

	function getJSIntactURL(module){
		return config.baseDir + module + config.jsFileTail;
	}

	function loadModule(module, callback){

		if(moduleMap[module]) return;

		var url = getJSIntactURL(module),
			s = d.createElement("script");
		s.type = "text/javascript";
		s.setAttribute("data-name", module);
		s.src = url;
		s.onload = function(e){
			var t = e.target,
				modName = t.getAttribute('data-name');
			moduleMap[modName] = moduleCache;
			callback(modName, moduleCache);
			moduleCache = null;		
		};
		d.body.appendChild(s);
	}

	function loadModules(modMap, callback){
		
		var modNum = 0,
			loadedNum = 0,
			m = {};

		for(var i in modMap) {
			modNum++;
			loadModule(modMap[i], function(k, v){
				loadedNum++;
				m[k] = v;
				if(loadedNum === modNum) callback(m);
			});
		}
	}

	
	/**/
	function getSelfElem(){
		var s = d.getElementsByTagName('script');
		return s[s.length-1];
	}

	w.require = function(deps, factory){

		if(typeof deps === 'function') {
			factory = deps;
			factory(runTime, {});
		} else {
			loadModules(deps, function(m){
				factory(runTime, m);
			})
		}
	};

	w.define = function(){
		
	};

	//init
	function init(){
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
		
	}
	init();

})(window);