
$(function () {
    $.ajax({
        url:"/videoCTL/list",
        success:function (data) {
            showData(data)
        }
    });

    $('#searchForm').ajaxForm({
        url:'/videoCTL/list',
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
            var $a=$('<a>').attr('href','/videoCTL/detail?id='+v.id).addClass('weui-media-box weui-media-box_appmsg');
            var $hd=$('<div>').addClass('weui-media-box__hd').append(
                $('<img>').addClass('weui-media-box__thumb').attr('src',v.poster));
            var $bd=$('<div>').addClass('weui-media-box__bd').append(
                $('<h4>').addClass('weui-media-box__title').html(v.name)).append(
                $('<p>').addClass('weui-media-box__desc').html(v.brief));
            $('.weui-panel__bd').append($a.append($hd).append($bd))
        })
    }
}