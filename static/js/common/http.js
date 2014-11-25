define(function(){

	

	function ajax(conf, onchange){
		
		var xhr = new XMLHttpRequest();
		xhr.open(conf.type, conf.url, conf.async);
		xhr.send(conf.data);
		xhr.onreadystatechange = onchange;
	}

	function get(url, success){
		ajax({
			url: url,
			type: 'get',
			async: true,
			data: null
		}, function(e){
			var t = e.target;
			if(t.status === 4 && t.state == 200){
				success(t.respanse);
			}
		});
	}

	function post(){

	}

	return {
		ajax: ajax,
		get: get,
		post: post
	};

});