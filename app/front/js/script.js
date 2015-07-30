var conf = {
    pictures_api_url: "",
    flag_column_count: 10,
    flag_row_count: 10,
    zoom_max: 8,
    zoom_min: 0,
    layout: 0
};

var init = function() {
    $("#button_in").click(function() {
        zoom_in();
    });

    $("#button_out").click(function() {
        zoom_out();
    });

    init_structure();
    createMap(conf.layout);
};

var createMap = function(layout) {
    for(var row=0;row<conf.flag_row_count;row++) {
        for(var column=0;column<conf.flag_row_count;column++) {
            var url = getPicture(layout, column, row);
            $('#tile_' + column + '_' + row).attr("src", url);
        }
    }
};

var getPicture = function(layout, x, y) {
    $.ajax({
        url: conf.pictures_api_url + '/' + layout + '/' + x + '_' + y + '.png'
    })
    .done(function( data ) {
        return data;
    });
};

var init_structure = function() {
    for (var row = 0; row < conf.flag_row_count; row++) {
        var section_name = 'section_'+row;
        $("#main").append("<div id=\'"+section_name+"\' class='section group'>");
        for (var column = 0; column < conf.flag_column_count; column++) {
            var span_name = 'span_' + column;
            var tile_name = 'tile_' + column + "_" + row;
            $("#" + section_name).append("<div id=\'" + section_name + '_' + span_name + "\' class='col span_1_of_10'>");
            $("#" + section_name+'_'+span_name).append("<img id=\'" + tile_name + "\'/>");
        }
        $("#main").append( "</div>" );
    }
};

var zoom_in = function() {
    if(conf.layout < conf.zoom_max) {
        conf.layout++;
        createMap();
    }
};

var zoom_out = function() {
    if(conf.layout > conf.zoom_min) {
        conf.layout--;
        createMap();
    }
};