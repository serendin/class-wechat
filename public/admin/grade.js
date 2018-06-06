
$(function () {

    $("#term").picker({
        title: "要查询的学期",
        cols: [
            {
                textAlign: 'center',
                values: ['2017/2018(2)', '2017/2018(1)', '2016/2017(2)', '2016/2017(1)',
                    '2015/2016(2)', '2015/2016(1)', '2014/2015(2)', '2014/2015(1)',
                    '2013/2014(2)', '2013/2014(1)', "2012/2013(2)", '2012/2013(1)'
                ]
            }
        ],
        onChange:function (p,v) {
            if(v.length>0){
                if(gradeList.hasOwnProperty(v[0])){
                    showData(gradeList[v[0]])
                }else{
                    $.ajax({
                        url:"/gradeCTL/list?term="+v[0],
                        success:function (data) {
                            showData(data);
                            gradeList[v[0]]=data;
                        }
                    });
                }

            }
        }
    });
});
var gradeList={};
function showData(data) {
    if(data!=undefined||data.length!=0){
        $('.weui-panel__bd').empty()
        $.each(data,function (i,v) {
            var $div=$('<div>').addClass('weui-media-box weui-media-box_text');
            var $h=$('<h4>').addClass('weui-media-box__title').html(v.lesson_name+" : "+v.score);
            var $ul=$('<ul>').addClass('weui-media-box__info').append(
                $('<li>').addClass('weui-media-box__info__meta').html(v.tea_name)).append(
                $('<li>').addClass('weui-media-box__info__meta').html("学分"+v.credit));
            $('.weui-panel__bd').append($div.append($h).append($ul))
        })
    }
}