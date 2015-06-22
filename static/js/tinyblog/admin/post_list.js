define({

	headerFn: 'tinyblog/admin/m_header',
	postPreView: 'tinyblog/admin/m_post_preview'

}, function(global, modules){

	"use strict"

	var headerFn = modules.headerFn;
	var postPreView = modules.postPreView;

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
				ballback();
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
		hideAllPenal: function () {
			postPreView.hide();
			headerFn.display(false);
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

			if (postNum == 1) {

				let currId = Object.keys(selPost)[0];
				let d = curPostData[currId];
				postPreView.show(d);
				headerFn.display(false);

			} else if(postNum > 1) {

				postPreView.hide();
				headerFn.display(true);

			} else {

				UI.hideAllPenal();

			}
		},
		unSelectAllPost: function () {
			selPost = {};
			UI.unSelectAllPost();
			UI.hideAllPenal();
		}
	};

	function eventSwitch (e) {

		var elem = e.target;
		var cList = elem.classList;

		if (cList.contains("j-itemopt")) {
			
			Fn.selectPostItem(elem);

		} else if (cList.contains("j-delpost")) {

			let id = elem.getAttribute("data-id");
			delPost(id, function () {
				refreshPostList();
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