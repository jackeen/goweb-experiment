define({

	//postFn: 'tinyblog/post_data'

}, function(global, modules){

	"use strict"

	const postListURL = "/api/postlist/get"

	function buildPostList(list) {
		
		var listStr = "", item = {};
		var len = list.length;

		for (var i = 0; i<len; i++) {
			item = list[i];
			listStr += `<li><h2>${item.title}</h2>\
							<p>${item.createTime}-${item.author}</p>\
						</li>`;
		}

		return listStr;
	}

	fetch(postListURL, {
		method: "get"
	}).then(function (res) {

		return res.json();

	}).then(function (d) {
 
		if (d.state) {
			let listCon = document.querySelector("#postlist");
			listCon.innerHTML = buildPostList(d.data);
			console.log("all post count: ", d.count);
		} else {
			return Error(d.message);
		}

	}).catch(function (err) {
		alert(err);
	});


});