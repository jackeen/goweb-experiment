define({

	//postFn: 'tinyblog/post_data'
	headerFn: 'tinyblog/admin/m_header'

}, function(global, modules){

	"use strict"

	var headerFn = modules.headerFn;

	const postListURL = "/api/postlist/get"

	var selPost = {};
	var curPostData = {};

	//the function for post list rander that base on ES6 string template
	function buildPostData(list) {
		
		var listStr = "", item = {};
		var len = list.length;
		var data = {};

		for (var i = 0; i<len; i++) {
			item = list[i];
			listStr += `\
				<li>\
					<h2 data-id="${item.id}" class="title">${item.title}</h2>\
					<p class="meta">${item.author} - ${item.createTime}</p>\
				</li>`;
			data[item.id] = item;
		}

		return {
			html: listStr,
			data: data
		};
	}



	//get post list
	function refreshPostList() {

		fetch(postListURL, {
			method: "get"
		}).then(function (res) {

			return res.json();

		}).then(function (d) {
	 
			if (d.state) {
				let listCon = document.querySelector("#postlist");
				let postData = buildPostData(d.data);
				listCon.innerHTML = postData.html;
				curPostData = postData.data;
				console.log("all post count: ", d.count);
			} else {
				return Error(d.message);
			}

		}).catch(function (err) {
			alert(err);
		});
	}


	function optPostInfo(id, isShow) {
		var con = document.querySelector("#postinfocon");
		if(isShow) con.classList.add("show");
		else con.classList.remove("show");
	}

	function optPostList(ids) {
		
	}

	
	function init() {

		refreshPostList();

		//post list event router
		document.querySelector("#postlist").onclick = function (e) {
			
			var t = e.target;
			var id = t.getAttribute("data-id");
			
			if (id !== "") {
				if (t.classList.contains("selected")) {
					t.classList.remove("selected");
					delete selPost[id];
				} else {
					t.classList.add("selected");
					selPost[id] = true;
				}
			}

			let ids = Object.keys(selPost);
			if (ids.length > 1) {
				optPostList(ids);
				headerFn.display(true);
				optPostInfo("", false);
			} else if(ids.length === 1) {
				optPostInfo(ids[0], true);
				headerFn.display(false);
			} else {
				optPostInfo("", false);
			}
			
		};
	}

	init();
});