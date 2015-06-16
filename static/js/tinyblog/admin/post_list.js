define({

	//postFn: 'tinyblog/post_data'
	headerFn: 'tinyblog/admin/m_header'

}, function(global, modules){

	"use strict"

	const postListURL = "/api/postlist/get"

	var seledPost = [];

	//the function for post list rander that base on ES6 string template
	function buildPostList(list) {
		
		var listStr = "", item = {};
		var len = list.length;

		for (var i = 0; i<len; i++) {
			item = list[i];
			listStr += `\
				<li>\
					<h2 class="title">\
						<input class="j-sel" type="checkbox" value="${item.id}">\
						${item.title}\
					</h2>\
					<p class="meta">${item.author} - ${item.createTime}</p>\
				</li>`;
		}

		return listStr;
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
				listCon.innerHTML = buildPostList(d.data);
				console.log("all post count: ", d.count);
			} else {
				return Error(d.message);
			}

		}).catch(function (err) {
			alert(err);
		});
	}


	function optPostInfo() {

	}

	function optPostList(isEditList) {
		var hFn = modules.headerFn;
		hFn.display(isEditList);
	}

	
	function init() {

		refreshPostList();

		//post list event router
		document.querySelector("#postlist").onclick = function (e) {
			
			var t = e.target;
			if (t && t.className == "j-sel") {
				let v = t.value;
				if (t.checked) {
					seledPost.push(v);
				} else {
					let i = seledPost.indexOf(v);
					if(i > -1) seledPost.splice(i, 1);
				}
			}

			if (seledPost.length > 1) {
				optPostList(true);
			} else {
				optPostList(false);
			}
		};
	}

	init();
});