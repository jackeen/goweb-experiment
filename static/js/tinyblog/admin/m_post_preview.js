define({
	postData: 'tinyblog/post_data'
}, function (global, modules) {

	var PD = modules.postData;

	var ModuleEvent = {
		onDelPost: function (id) {},
		onCancel: function () {}
	};

	var con = document.createElement("div");
	con.id = "postinfocon";
	con.className = "post-info-con";
	document.body.appendChild(con);

	con.onclick = function (e) {
		e.stopPropagation();
		var t = e.target;
		var opt = t.getAttribute("data-opt");
		switch (opt) {
			case "del" :
				delTarget(t.getAttribute("data-id"));
				break;
			case "cancel" :
				hide();
				ModuleEvent.onCancel();
				break;
		}
	}

	function delTarget(id) {
		PD.delPost(id, function () {
			hide();
			ModuleEvent.onDelPost(id);
		}, function (err) {
			alert(err);
		});
	}
	
	function buildDom(d) {
		return `\
			<p class="p-cate p-item">${d.cate}</p>\
			<h1 class="p-title p-item">${d.title}</h1>\
			<p class="btns p-item">\
				<a data-opt="cancel" href="javascript:;">cancel</a>\
				<a data-opt="del" data-id="${d.id}" href="javascript:;">del</a>\
				<a href="/admin/editpost?id=${d.id}" target="_blank">edit</a>\
				<a href="/post/${d.title}" target="_blank">veiw</a>\
			</p>\
			<div class="p-content p-item">${d.content}</div>\
			<p>${d.tags.join()}</p>\
			<p class="p-author-date">\
				<a href="###">${d.author}</a>\
				-\
				<a href="###">${d.createTime}</a>\
			</p>\
		`;
	}

	function show(d) {
		con.innerHTML = buildDom(d);
		con.classList.add("show");
	}

	function hide() {
		con.classList.remove("show");
	}

	return {
		Event: ModuleEvent,
		show: show,
		hide: hide
	};

});