define({

	postFn: 'tinyblog/post_data'

}, function(global, modules){

	var postFn = modules.postFn;
	
	var title = document.getElementById('title');
	var content = document.getElementById('content');
	var saveBtn = document.getElementById('savebtn');

	saveBtn.onclick = function () {
		postFn.savePost({
			title: title.value,
			content: content.value
		}, function (res) {
			console.log(res);
		});
	};

});