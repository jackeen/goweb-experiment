define(function(global, modules){

	"use strict"

	function buildTagList(list) {
		var str = '';
		var item = null;
		for (var i = 0, len = list.length; i < len; i++) {
			item = list[i];
			str += `\
				<li>\
					<span data-opt="tagsel" class="name">${item.name}</span>\
					<span data-opt="tagdel" class="del" data-name="${item.name}">Del</span>\
				</li>\
			`;
		}
		return str;

	}

	function delTag(name, resolve, reject) {

		var url = "/api/tag/del" + "?n=" + name;

		fetch(url, {
			credentials: "same-origin"
		}).then(function (res) {

			return res.json();

		}).then(function (d) {
	 
			if (d.state) {
				resolve(d)
			} else {
				reject();
			}

		}).catch(function (err) {
			alert(err);
		});
	}

	function init() {

		var tagList = document.querySelector("#taglist");
		var createBtn = document.querySelector("#tagcreatebtn");
		var ipt = document.querySelector("#tagnamein");

		tagList.onclick = function (e) {
			var t = e.target;
			var opt = t.getAttribute("data-opt");

			if (opt === "tagdel") {
				let name = t.getAttribute("data-name");
				delTag(name, function () {

				}, function () {

				});
			}
		};

		createBtn.onclick = function () {
			var v = ipt.value;
		};

		ipt.onblur = function () {
			var v = this.value;
		}



		fetch("/api/tag/get", {
			credentials: "same-origin"
		}).then(function (res) {

			return res.json();

		}).then(function (d) {
	 
			if (d.state) {
				tagList.innerHTML = buildTagList(d.data);
			} else {
				return Error(d.message);
			}

		}).catch(function (err) {
			alert(err);
		});

	}

	init();

});