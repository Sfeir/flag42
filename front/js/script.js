var init = function() {

};

var init_structure = function() {
    //creation de la structure
    for (var line = 0; line < 10; line++) {
        var section_name = 'section_'+line;
        $("#main").append("<div id=\'"+section_name+"\' class='section group'>");
        for (var colomn = 0; colomn < 10; colomn++) {
            var span_name = 'span_'+colomn;
            $("#"+section_name).append("<div id=\'"+section_name+'_'+span_name+"\' class='col span_1_of_10'>");
        }
        $("#main").append( "</div>" );
    }
};

var zoom_plus = function() {
    alert('+');
};

var zoom_moins = function() {
    alert('-');
};
