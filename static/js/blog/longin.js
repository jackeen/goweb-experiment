define({

	http: 'common/http',
	cookie: 'common/cookie',
	animation: 'common/animation'

}, function(global, modules){

	console.log('login: ', modules);

	var http = modules.http;
	
	var f = document.forms['login'],
		btn = f['submit'],
		url = f.action;

	btn.onclick = function(e) {

		var f = this.form;

		http.post(url, {
			user: f['user'].value,
			pass: f['pass'].value
		}, function(res) {
			console.log(res);
		});
	};

});