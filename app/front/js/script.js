var conf = {
    pictures_api_url: "https://storage.googleapis.com/flag-42/eiffeltower",
    flag_column_count: 10,
    flag_row_count: 10,
    position_x: 0,
    position_y: 0,
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

    $("#button_left").click(function() {
        move_left();
    });

    $("#button_right").click(function() {
        move_right();
    });

    $("#button_up").click(function() {
        move_up();
    });

    $("#button_down").click(function() {
        move_down();
    });

    init_structure();
    createMap(conf.layout);
};

var createMap = function(layout) {
    for(var row=0;row<conf.flag_row_count * Math.pow(2,conf.layout);row++) {
        for(var column=0;column<conf.flag_row_count * Math.pow(2,conf.layout);column++) {
            var url = getPicture(layout, column + conf.position_x, row + conf.position_y);
            $('#tile_' + column + '_' + row).attr("src", url);
        }
    }
};

var getPicture = function(layout, x, y) {
    return conf.pictures_api_url + '/' + layout + '/' + x + '_' + y + '.png';
};

var init_structure = function() {
    $("#main").empty();
    for (var row = 0; row < conf.flag_row_count * Math.pow(2,conf.layout); row++) {
        var section_name = 'section_'+row;
        $("#main").append("<div id=\'"+section_name+"\' class='section group'>");
        for (var column = 0; column < conf.flag_column_count * Math.pow(2,conf.layout); column++) {
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
        init_structure();
        createMap(conf.layout);
    }
};

var zoom_out = function() {
    if(conf.layout > conf.zoom_min) {
        conf.layout--;
        init_structure();
        createMap(conf.layout);
    }
};

var move_left = function() {
    conf.position_x--;
    createMap(conf.layout);
};

var move_right = function() {
    conf.position_x++;
    createMap(conf.layout);
};

var move_up = function() {
    conf.position_y--;
    createMap(conf.layout);
};

var move_down = function() {
    conf.position_y++;
    createMap(conf.layout);
};
