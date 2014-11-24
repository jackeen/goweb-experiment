define(function(){

	

	function ajax(conf, onchange){
		
		var xhr = new XMLHttpRequest();
		xhr.open(conf.type, conf.url, conf.async);
		xhr.send(conf.data);
		xhr.onreadystatechange = onchange;
		return xhr;
	}

	function get(){
		var xhr = ajax({
			url: url,
			type: 'get',
			async: true
		}, function(e){
			var t = e.target;
			if(t.status === 4 && t.state == 200){
				
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