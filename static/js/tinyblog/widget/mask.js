define(function(G){

	"use strict"

	G.loadStyle('tinyblog/widget');

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