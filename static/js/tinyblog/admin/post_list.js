define({

	//postFn: 'tinyblog/post_data'

}, function(global, modules){

	fetch("/api/postlist/get").then(function (res){
		console.log(res);
		window.res = res;
	}, function (e){
		console.log(e);
	});

});