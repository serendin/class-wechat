
$(function () {
    $.ajax({
        url:"/questionCTL/list",
        success:function (data) {
            showData(data)
        }
    });

    $.ajax({
        url:"/questionCTL/myList",
        success:function (data) {
            showMyData(data)
        }
    });

    $('#searchAll').ajaxForm({
        url:'/questionCTL/list',
        success:function(data){
            $('.question_list').empty();
            showData(data);
        },
        error:function(){
            $.toast("信息获取失败","cancel")
        }
    }).submit(function () {
        return false;
    });

    $('#searchMine').ajaxForm({
        url:'/questionCTL/myList',
        success:function(data){
            $('.question_mine').empty();
            showData(data);
        },
        error:function(){
            $.toast("信息获取失败","cancel")
        }
    }).submit(function () {
        return false;
    })
});

function showData(data) {
    if(data!=undefined||data.length!=0){
        $.each(data,function (i,v) {
            var $div=$('<div>').addClass('weui-media-box weui-media-box_text');
            var $h=$('<h4>').addClass('weui-media-box__title').html(v.title).css('work-break','break-all');
            var $p=$('<p>').addClass('weui-media-box__desc').html(v.question);
            var $ul=$('<ul>').addClass('weui-media-box__info').append(
                $('<li>').addClass('weui-media-box__info__meta').html(v.stu_name)).append(
                $('<li>').addClass('weui-media-box__info__meta  weui-media-box__info__meta_extra').html(v.lesson_name)).append(
                $('<li>').addClass('weui-media-box__info__meta  weui-media-box__info__meta_extra').html(v.updated_at.substr(5,11).replace('T',' ')));
            $('.question_list').append($div.append($h).append($p).append($ul))

        })
    }
}

function showMyData(data) {
    if(data!=undefined||data.length!=0){
        $.each(data,function (i,v) {
            var $div=$('<div>').addClass('weui-media-box weui-media-box_text');
            var $h=$('<h4>').addClass('weui-media-box__title').html(v.title).css('work-break','break-all');
            var $p=$('<p>').addClass('weui-media-box__desc').html(v.question);
            var $ul=$('<ul>').addClass('weui-media-box__info').append(
                $('<li>').addClass('weui-media-box__info__meta').html(v.stu_name)).append(
                $('<li>').addClass('weui-media-box__info__meta  weui-media-box__info__meta_extra').html(v.lesson_name)).append(
                $('<li>').addClass('weui-media-box__info__meta  weui-media-box__info__meta_extra').html(v.updated_at.substr(5,11).replace('T',' ')));
            $('.question_mine').append($div.append($h).append($p).append($ul))

        })
    }
}