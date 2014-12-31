define(function(global, modules){

	function getAjaxData(obj) {
		var s = '';
		for(var i in obj) {
			s += '&' + i + '=' + obj[i];
		}
		return s.replace('&', '');
	}

	function ajax(conf, onchange) {
		
		var xhr = new XMLHttpRequest(),
			data = null;

		if(conf.data) {
			if(conf.dataType === 'json') {
				data = JSON.stringify(conf.data);
			} else {
				data = getAjaxData(conf.data);
			}
		}

		xhr.open(conf.type, conf.url, conf.async);

		if (conf.dataType !== 'json') {
			xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
		}

		xhr.send(data);
		xhr.onreadystatechange = onchange;
	}

	function get(url, data, success) {
		ajax({
			url: url,
			type: 'get',
			async: true,
			data: null
		}, function(e) {

			if(this.readyState === 4 && this.status === 200){
				success(this.responseText);
			}
		});
	}

	function post(url, data, success) {
		ajax({
			url: url,
			type: 'post',
			async: true,
			data: data
		}, function(e) {

			if(this.readyState === 4 && this.status === 200) {
				success(this.responseText);
			}
		});
	}

	return {
		ajax: ajax,
		get: get,
		post: post
	};

});