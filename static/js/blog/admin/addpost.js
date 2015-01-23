define({

	http: 'lib/http'

}, function(global, modules){

	var http = modules.http;
	
	var f = document.forms['addpost'],
		btn = f['save'],
		url = f.action;

	btn.onclick = function(e) {

		var f = this.form;

		http.post(url, {

			title: f['title'].value,
			content: f['content'].value,
			//allowcomment: 'allow',
			draft: 'xx'

		}, function(res) {

			console.log(res);
			
		});
	};

});