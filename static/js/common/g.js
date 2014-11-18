//common base html5

(function(w, globalName){

	var d = w.document;

	var config = {
		baseDir : '/js/',
		jsFileTail : ".js",
		cssFileTail : ".css"
	};

	var scriptMap = {};

	function loadScript(url, fn){

		if (scriptMap[url]) {
			fn();
			return;
		};

		s = d.createElement("script");
		s.type = "text/javascript";
		s.src = url;
		s.onload = function(){
			fn();
		};
		d.body.appendChild(s);
		scriptMap[url] = true;
	}

	function loadScriptList(){

	}

	function loadCss(url){
		
	}

	function getJSIntactURL(modName){
		return config.baseDir + modName + config.jsFileTail;
	}

	function getCSSIntactURL(modName){
		return config.baseDir + modName + config.cssFileTail;
	}

	var $ = {
		
		config: function(){

		},

		require: function(){

		},

		define: function(jsList, fn){

		}

	};

	w[globalName] = $;

})(window, "$");