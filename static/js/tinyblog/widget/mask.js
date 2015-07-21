define(function(G, M){

	"use strict"

	G.loadStyle('widget/base');

	function getMask() {
		var m = document.createElement('div');
		m.className = 'm-mask';
		document.body.appendChild(m);
		return m;
	}

	return {
		mask: getMask
	};

});