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
					<div class="j-itemopt item-opt">#<div class="j-itembtns item-btns">\
						<a data-id="${item.id}" class="j-selpost" href="javascript:;">select</a>\
						<a data-id="${item.id}" class="j-delpost" href="javascript:;">del</a>\
						<a href="/post/${item.title}" target="_blank">outveiw</a>\
					</div></div>\
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

	function delPost(id, ballback) {

		var url = delPostURL + "?id=" + id;

		fetch(url).then(function (res) {
			return res.json();
		}).then(function (d) {

			if (d.state) {
				classList();
			}

		}).catch(function (err) {
			alert(err);
		});

	}


	var UI = {
		hideAllItemBtn: function () {
			var btns = document.querySelectorAll("#postlist .j-itembtns.show");
			btns = Array.prototype.slice.call(btns);
			btns.forEach(function (elem) {
				elem.classList.remove("show");
			});
		},
		displayItemBtns: function (elem) {
			var btnCon = elem.querySelector(".j-itembtns");
			btnCon.classList.toggle("show");
		},
		displayPostInfo: function (isShow) {
			var con = document.querySelector("#postinfocon");
			if(isShow) con.classList.add("show");
			else con.classList.remove("show");
		}
	};

	function eventSwitch (elem) {
		
		var cList = elem.classList;

		if(cList.contains("j-itemopt")) {
			
			UI.hideAllItemBtn();
			UI.displayItemBtns(elem);

		} else if(cList.contains("j-selpost")) {

			let id = elem.getAttribute("data-id");
			//selPost[id] = true;
			UI.displayPostInfo(true);

		} else if(cList.contains("j-delpost")) {

			let id = elem.getAttribute("data-id");
			delPost(id, function () {

			});

		} else {
			UI.hideAllItemBtn();
		}
	}

	
	function init() {

		refreshPostList();

		document.onclick = function (e) {
			var t = e.target;
			eventSwitch(t);
		}

	}

	init();
});