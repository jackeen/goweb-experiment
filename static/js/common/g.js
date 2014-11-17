//common base html5

(function(w, globalName){

	var d = w.document;

	var scriptMap = {};

	function loadScriptList(){

	}

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