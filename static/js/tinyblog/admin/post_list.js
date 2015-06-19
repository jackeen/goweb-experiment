define({

	//postFn: 'tinyblog/post_data'
	headerFn: 'tinyblog/admin/m_header'

}, function(global, modules){

	"use strict"

	var headerFn = modules.headerFn;

	const postListURL = "/api/postlist/get"
	const delPostURL = "/api/post/del";

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
					<div data-id="${item.id}" class="j-itemopt item-opt"></div>\
					<h2 class="title">${item.title}</h2>\
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
			credentials: "same-origin"
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

	function delPost(id, ballback) {

		var url = delPostURL + "?id=" + id;

		var req = new Request(url, {
			credentials: "same-origin"
		});

		fetch(req).then(function (res) {
			return res.json();
		}).then(function (d) {

			if (d.state) {
				refreshPostList();
			}

		}).catch(function (err) {
			alert(err);
		});

	}


	var UI = {
		selectPostItem: function (elem) {
			elem.classList.toggle("selected");
		},
		unSelectAllPost: function () {
			var postList = document.querySelectorAll("#postlist .j-itemopt.selected");
			postList = Array.prototype.slice.call(postList);
			postList.forEach(function (elem) {
				elem.classList.remove("selected");
			});
		},
		displayPostInfo: function (isShow) {
			var con = document.querySelector("#postinfocon");
			if(isShow) con.classList.add("show");
			else con.classList.remove("show");
		},
		changeOptPanel: function (postNum) {
			if (postNum == 1) {
				UI.displayPostInfo(true);
				headerFn.display(false);
			} else if(postNum > 1) {
				UI.displayPostInfo(false);
				headerFn.display(true);
			} else {
				UI.displayPostInfo(false);
				headerFn.display(false);
			}
		}
	};

	var Fn = {
		selectPostItem: function (elem) {
			var id = elem.getAttribute("data-id");
			if(selPost[id]) {
				delete selPost[id];
			} else {
				selPost[id] = true;
			}
			UI.selectPostItem(elem);

			var postNum = Object.keys(selPost).length;
			UI.changeOptPanel(postNum);
		},
		unSelectAllPost: function () {
			selPost = {};
			UI.unSelectAllPost();
			UI.changeOptPanel(0);
		}
	};

	function eventSwitch (e) {

		var elem = e.target;
		var cList = elem.classList;

		if(cList.contains("j-itemopt")) {
			
			Fn.selectPostItem(elem);

		} else if(cList.contains("j-selpost")) {

			let id = elem.getAttribute("data-id");
			//selPost[id] = true;
			UI.displayPostInfo(true);

		} else if(cList.contains("j-delpost")) {

			let id = elem.getAttribute("data-id");
			delPost(id, function () {

			});

		} else {
			Fn.unSelectAllPost();
		}
	}

	
	function init() {

		refreshPostList();

		document.onclick = function (e) {
			eventSwitch(e);
		}

	}

	init();
});