/**
 * Created by JUNYI on 2017-08-30.
 */
function init() {
    var map = L.map('map', {
        renderer: L.canvas.screen()
    }).setView([37.8, -96], 5);
    // add basemap
    L.tileLayer('https://api.mapbox.com/styles/v1/mapbox/navigation-guidance-night-v2/tiles/256/{z}/{x}/{y}?access_token=your_token', {
        attribution: '&copy; <a href="http://mapbox.com">Mapbox</a> contributors',
        maxZoom: 18
    }).addTo(map);

    function onEachFeature(feature, layer) {
        if (feature.properties) {
            var htmlStr = "<div><table>";
            var imageRow = "";
            for (key in feature.properties) {
                var rowStr = "";
                if (key == "ImageURL") {
                    imageRow = "<tr><td>" + key + "</td><td><a target=\"_blank\" href=\"" + feature.properties[key] + "\"><img src=" + feature.properties[key] + " style=\"width:120px; height:100px;\"/></a></td></tr>";
                    continue;
                }
                rowStr = "<tr><td>" + key + "</td><td>" + feature.properties[key] + "</td></tr>";
                htmlStr += rowStr;
            }
            htmlStr += imageRow;
            htmlStr += "</table></div>";
            layer.bindPopup(htmlStr);
        }
    }

    //add flickr points
    // from local file
    // add geojson as circle marker
    var markerOptions = {
        radius: 3,
        fillColor: "#ff7800",
        color: "#000",
        weight: 0,
        opacity: 0.5,
        fillOpacity: 0.5
    };

    L.geoJSON(featureCollections, {
        pointToLayer: function (feature, latlng) {
            // var latlng = L.latlng(feature.)
            return L.circleMarker(latlng, markerOptions);
        },
        onEachFeature: onEachFeature
    }).addTo(map);

    // from flickr server
    // $.get("http://localhost:8000/flickr_points", function (data) {
    //     var featureCollections = JSON.parse(data);

    //     // add geojson as circle marker
    //     var markerOptions = {
    //         radius: 3,
    //         fillColor: "#ff7800",
    //         color: "#000",
    //         weight: 0,
    //         opacity: 0.5,
    //         fillOpacity: 0.5
    //     };

    //     L.geoJSON(featureCollections, {
    //         pointToLayer: function (feature, latlng) {
    //             // var latlng = L.latlng(feature.)
    //             return L.circleMarker(latlng, markerOptions);
    //         },
    //         onEachFeature: onEachFeature
    //     }).addTo(map);
    // })
    //     .fail(function () {
    //         alert("error");
    //     });
}
