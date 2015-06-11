define({

	http: 'lib/http'

}, function(global, modules){

	var http = modules.http;

	var saveURL = '/api/post/put';

	function savePost(data, fn) {

		http.post(saveURL, {

			title: data.title,
			content: data.content,
			allowcomment: data.allowComment,
			draft: data.isDraft

		}, fn);
	}

	return {
		savePost: savePost
	};

});