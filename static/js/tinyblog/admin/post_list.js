define({

	//postFn: 'tinyblog/post_data'

}, function(global, modules){

	const postListURL = "/api/postlist/get"

	function buildPostList(list) {

	}

	fetch(postListURL, {
		method: "get"
	}).then(function (res) {

		return res.json();

	}, function (err) {

		console.log("network:", err);

	}).then(function (d) {

		if (d.state) {
			buildPostList(d.data)
			console.log("all post count: ", d.count);
		} else {
			console.log(d.message);
		}

	}).catch(function (err) {
		console.log(err);
	});

});