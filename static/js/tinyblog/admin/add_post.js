define({

	postFn: 'tinyblog/post_data',
	tag: 'tinyblog/admin/m_post_tag'

}, function(G){

	"use strict"

	var postFn = G.require('postFn');
	var Tag = G.require('tag');
	
	var title = document.getElementById('title');
	var content = document.getElementById('content');
	var saveBtn = document.getElementById('savebtn');

	function clickPSwitch() {
		this.parentNode.classList.toggle("off");
	}

	function switchEvent() {
		var s = document.querySelectorAll(".j-p-switch");
		s = Array.prototype.slice.call(s);
		s.forEach(function (elem, i) {
			elem.onclick = clickPSwitch;
		});
	}

	switchEvent();

	function savePost() {

		var f = new FormData(document.forms["postform"]);

		var req = new Request("/api/post/put", {
			credentials: "same-origin",
			method: "post",
			headers: {
 			   "Content-Type": "application/x-www-form-urlencoded"
 			},
			body: f
		});

		fetch(req).then(function (res) {
			return res.json();
		}).then(function (d) {
			

		}).catch(function (err) {
			
		});
	}

	saveBtn.onclick = function () {

		savePost();
/*		var f = new FormData(document.forms["postform"]);
		var xhr = new XMLHttpRequest();
		xhr.open("post", "/api/post/put");
		xhr.onload = function (e) {
			var t = e.target;
			if (t.status === 200) {
				//t.response
			}
		}
		xhr.send(f);*/
	};

});