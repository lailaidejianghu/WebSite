
    function userChange(x) {
        var user = document.getElementById(x).value;
        //验证手机规则
        var regPh = /^1[3|4|5|7|8][0-9]{9}$/;
        //验证邮箱规则
        var regEml=/^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/;
        
        if (!regPh.test(user) && !regEml.test(user)) {
            document.getElementById(x).style="color:red;border-color:red";
            document.getElementById(x).value="账号需手机或邮箱";
           
        }
     }

     function pwdChange(x) {
        var pwd = document.getElementById(x).value;
        if ((pwd.length)<6) {
            document.getElementById(x).style="color:red;border-color:red";
            document.getElementById(x).value="密码长度不少于6！";
           
        }
     }

    function inputOnfocus(x) {
        document.getElementById(x).style="color:black;border-color:darkgray";
        document.getElementById(x).value="";
    }
