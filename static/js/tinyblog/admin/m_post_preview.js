define(function (global, modules) {

	var con = document.createElement("div");
	con.id = "postinfocon";
	con.className = "post-info-con";
	document.body.appendChild(con);

	con.onclick = function (e) {
		e.stopPropagation();
	}
	
	function buildDom(d) {
		return `\
			<h1 class="p-title">${d.title}</h1>\
			<p class="p-cate">${d.cate}</p>\
			<div class="p-content">${d.content}</div>\
			<p class="p-author-date">\
				<a href="###">${d.author}</a>\
				-\
				<a href="###">${d.createTime}</a>\
			</p>\
			<p class="btns">\
				<a data-id="${d.id}" href="javascript:;">edit</a>\
				<a data-id="${d.id}" href="javascript:;">del</a>\
				<a href="/post/${d.title}" target="_blank">veiw</a>\
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
		show: show,
		hide: hide
	};

});