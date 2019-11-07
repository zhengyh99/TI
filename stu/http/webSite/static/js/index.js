function showme(){
    alert("Hi,I'm Victor,I love LJX!!!")
}

function showOtherUser(){
    $.post("/userlist",function(data){
        var result = "";
        for(i=0;i<data.length;i++){
            result += "<tr>";
            result += "<td>"+data[i].uname+"</td>";
            result += "<td>"+data[i].age+"</td>";
            result += "</tr>";
        }
        $("#ulist").html(result);
    })
}