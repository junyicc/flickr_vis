/**
 * Created by JUNYI on 2017-08-30.
 */
function init() {
    var map = L.map('map').setView([37.8, -96], 5);
    L.tileLayer('https://api.mapbox.com/styles/v1/mapbox/navigation-guidance-night-v2/tiles/256/{z}/{x}/{y}?access_token=pk.eyJ1IjoianVueWktY2FpIiwiYSI6ImNqMHEybDlrdjAwdWIyeW81bXRxMW5uZzYifQ.mJij_8lY1Exj7E4pfjJ7Ew', {
        attribution: '&copy; <a href="http://mapbox.com">Mapbox</a> contributors',
        maxZoom: 18
    }).addTo(map);

    function onEachFeature(feature, layer) {
        if (feature.properties) {
            layer.bindPopup("<img src=" + feature.properties.ImageURL + " alt=" + feature.properties.Description + " style=\"width:120px; height:100px;\"/>");
        }
    }

    //add flickr points
    var max, scale,
        classes = 5,
        scheme = colorbrewer["YlOrRd"][classes];
    /**
        * Hexbin style callback.
        *
        * Determines a quantize scale (http://bl.ocks.org/4060606) based on the
        * map's initial data density (which is based on the initial zoom level)
        * and applies a colorbrewer (http://colorbrewer2.org/) colour scheme
        * accordingly.
        */
    function hex_style(hexagons) {
        // Maintain a density scale relative to initial zoom level.
        if (!scale) {
            // quantize
            // max = d3.max(hexagons.data(), function (d) { return d.length; });
            // scale = d3.scale.quantize()
            //     .domain([0, max])
            //     .range(d3.range(classes));

            // scale cluster
            var hexagonsLen = []
            hexagons.data().forEach(function (elem) {
                hexagonsLen.push(elem.length)
            })

            scale = d3.scaleCluster()
                .domain(hexagonsLen.sort())
                .range(d3.range(classes))
        }

        hexagons
            .attr("stroke", scheme[classes - 1])
            .attr("fill", function (d) {
                return scheme[scale(d.length)];
            });
    }
    // from local file
    // L.hexLayer(featureCollections, {
    //     applyStyle: hex_style
    // }).addTo(map);

    // from flickr server
    d3.json("http://localhost:8000/flickr_points", function (err, data) {
        // When data arrives, create leaflet layer with custom style callback.
        if (err) return console.warn(err);
        if (data != null) {
            var featureCollections = JSON.parse(data);
            L.hexLayer(featureCollections, {
                applyStyle: hex_style
            }).addTo(map);
        }
    });
}
