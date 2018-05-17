
$(function () {
    $.ajax({
        url:"/noticeCTL/list",
        success:function (data) {
            showData(data);
        }
    });

    $('#searchForm').ajaxForm({
        url:'/noticeCTL/list',
        success:function(data){
            $('.weui-panel__bd').empty();
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
            var $h=$('<h4>').addClass('weui-media-box__title').html(v.lesson_name+":"+v.title);
            var $p=$('<p>').addClass('weui-media-box__desc').html(v.content);
            var $ul=$('<ul>').addClass('weui-media-box__info').append(
                $('<li>').addClass('weui-media-box__info__meta').html(v.tea_name)).append(
                $('<li>').addClass('weui-media-box__info__meta').html(v.updated_at.substr(5,11).replace('T',' ')));
            $('.weui-panel__bd').append($div.append($h).append($p).append($ul))
        })
    }
}