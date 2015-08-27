define({

	Mask: 'tinyblog/widget/mask'

}, function (G) {

	"use strict";

	function buildTagList(list) {
		var str = '';
		var item = null;
		for (var i = 0, len = list.length; i < len; i++) {
			item = list[i];			
			str += `\
				<li id="tagitem_${item.name}">\
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

	function saveTag(name, resolve, reject) {

		var url = "/api/tag/put" + "?n=" + name;

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

	

	var tagList = document.querySelector("#taglist");
	var createBtn = document.querySelector("#tagcreatebtn");
	var ipt = document.querySelector("#tagnamein");

	tagList.onclick = function (e) {
		var t = e.target;
		var opt = t.getAttribute("data-opt");

		if (opt === "tagdel") {
			let name = t.getAttribute("data-name");
			delTag(name, function () {

				var tag = document.querySelector("#tagitem_" + name);
				tag.remove();
				tag = null;

			}, function () {

			});
		} else if (opt === "tagsel") {
			t.classList.toggle("selected");
		}
	};

	createBtn.onclick = function () {
		var v = ipt.value;
		saveTag(v, function () {

			var dom = buildTagList([{name:v}]);
			var html = tagList.innerHTML;
			tagList.innerHTML = dom + html;
			ipt.value = "";

		}, function () {

		});
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

	function getTags() {
		var tags = tagList.querySelectorAll(".selected")
		var tagArr = Array.prototype.slice.call(tags);
		var tagNameArr = [];
		tagArr.forEach(function (v, i, self) {
			tagNameArr.push(v.innerText || v.contentText);
		});
		return tagNameArr;
	}

	

	return {
		getTags: getTags
	};

});