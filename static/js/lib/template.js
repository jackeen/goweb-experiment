define(function (global) {

	function parse(tpl, data) {
		return tpl.replace(/{{\w+}}/g/, function (v) {
			v.replace('{{', '').replace('}}', '');
			return data[v] + '';
		});
	}

	return {
		parse: parse
	};

})