# -*- coding: utf-8 -*-
import flickrapi
import psycopg2


def connect_db(**kw):
    conn = None
    try:
        conn = psycopg2.connect(**kw)
    except Exception as ex:
        print("failed to connect to the database, ", ex)
    return conn


def get_photos():
    # flickr info
    # add your key and secret here
    api_key = ""
    api_secret = ""

    flickr = None
    photos = None
    try:
        flickr = flickrapi.FlickrAPI(api_key, api_secret)
        photos = flickr.walk(bbox='-125, 25, -64, 48', accuracy=3,
                             extras='date_taken, geo, url_c, owner_name')
    except Exception as ex:
        print('Error: ', ex)
    # get photo url from photos
    if photos:
        # store into postgresql
        db_info = {
            'database': 'database',
            'user': 'username',
            'password': 'password',
            'host': 'host',
            'port': 'port'
        }
        conn = connect_db(**db_info)
        cur = conn.cursor()
        image_source = 'flickr'
        image_ids = []
        try:
            for photo in photos:
                image_id = photo.get('id')
                if image_id in image_ids:
                    continue
                image_ids.append(image_id)
                lat = photo.get('latitude')
                lon = photo.get('longitude')
                location = 'POINT({0} {1})'.format(lon, lat)
                image_url = photo.get('url_c')
                width = photo.get('width_c')
                height = photo.get('height_c')

                # new added
                owner = photo.get('owner')
                tag = photo.tag
                description = photo.get('title')
                taken_time = photo.get('datetaken')
                if location and image_url and width and height:
                    cur.execute(
                        """INSERT INTO "image_info" (image_id,
                        image_source, width, height, tags, image_url,
                        taken_time, owner, description, lat, lon, location)
                        SELECT %(image_id)s, %(image_source)s,
                        %(width)s, %(height)s, %(tag)s, %(image_url)s,
                        %(taken_time)s, %(owner)s, %(description)s,
                        %(lat)s, %(lon)s,
                        ST_GeomFromText(%(location)s, 4326)
                        WHERE NOT EXISTS(
                            SELECT image_id FROM "image_info" WHERE image_id = %(p_image_id)s
                        );""", {
                            'image_id': image_id,
                            'image_source': image_source,
                            'width': width,
                            'height': height,
                            'tag': tag,
                            'image_url': image_url,
                            'taken_time': taken_time,
                            'owner': owner,
                            'description': description,
                            'lat': lat,
                            'lon': lon,
                            'location': location,
                            'p_image_id': image_id,
                        })
                    conn.commit()
        except Exception as ex:
            print('Error: ', ex)
        finally:
            cur.close()
            conn.close()


if __name__ == '__main__':
    get_photos()
