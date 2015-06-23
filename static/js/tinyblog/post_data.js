define({

	http: 'lib/http'

}, function(global, modules){

	var http = modules.http;

	const saveURL = '/api/post/put';
	const delURL = "/api/post/del";

	function savePost(data, fn) {

		http.post(saveURL, {

			title: data.title,
			content: data.content,
			allowcomment: data.allowComment,
			draft: data.isDraft

		}, fn);
	}

	function delPost(id, resolve, reject) {

		var url = delURL + "?id=" + id;

		var req = new Request(url, {
			credentials: "same-origin"
		});

		fetch(req).then(function (res) {
			return res.json();
		}).then(function (d) {

			if (d.state) {
				resolve(d);
			}

		}).catch(function (err) {
			reject(err);
		});

	}

	return {
		savePost: savePost,
		delPost: delPost
	};

});