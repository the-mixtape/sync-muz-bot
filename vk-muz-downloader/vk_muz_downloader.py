import os
import m3u8
import requests
from vk_api import VkApi

from vk_api.audio import VkAudio
from Crypto.Cipher import AES
from Crypto.Util.Padding import unpad

APP_ID = 2685278


class VkMuzDownloader:
    def __init__(self, login: str, password: str):
        self.__vk_session = VkApi(
            login=login,
            password=password,
            app_id=APP_ID
        )
        self.__vk_session.auth()

        self.__vk_audio = VkAudio(self.__vk_session)

    def download_audios(self, temp_dir, out_dir, user_id):
        for audio_info in self.__vk_audio.get(owner_id=user_id):
            url = audio_info["url"]
            name = f'{audio_info["artist"]} {audio_info["title"]}'

            segments = self.__get_audio_segments(url=url)
            segments_data = self.__parse_segments(segments=segments)
            segments = self.__download_segments(segments_data=segments_data, index_url=url)

            self._convert_ts_to_mp3(segments=segments, out_name=name, out_dir=out_dir, tmp_dir=temp_dir)

    @staticmethod
    def _convert_ts_to_mp3(segments: bytes, out_name, out_dir, tmp_dir):
        tmp_path = os.path.join(tmp_dir, f'{out_name}.ts')
        out_path = os.path.join(out_dir, f'{out_name}.wav')
        with open(tmp_path, 'w+b') as f:
            f.write(segments)
        os.system(f'ffmpeg -i "{tmp_path}" -vcodec copy '
                  f'-acodec copy -vbsf h264_mp4toannexb "{out_path}"')
        os.remove(tmp_path)

    @staticmethod
    def __get_audio_segments(url: str):
        m3u8_data = m3u8.load(
            uri=url
        )
        return m3u8_data.segments

    @staticmethod
    def __parse_segments(segments: list):
        segments_data = {}

        for segment in segments:
            segment_uri = segment.uri

            extended_segment = {
                "segment_method": None,
                "method_uri": None
            }
            if segment.key.method == "AES-128":
                extended_segment["segment_method"] = True
                extended_segment["method_uri"] = segment.key.uri
            segments_data[segment_uri] = extended_segment
        return segments_data

    @staticmethod
    def __download_segments(segments_data: dict, index_url: str) -> bin:
        downloaded_segments = []

        for uri in segments_data.keys():
            audio = requests.get(url=index_url.replace("index.m3u8", uri))

            downloaded_segments.append(audio.content)

            if segments_data.get(uri).get("segment_method") is not None:
                key_uri = segments_data.get(uri).get("method_uri")
                key = VkMuzDownloader.__download_key(key_uri=key_uri)

                iv = downloaded_segments[-1][0:16]
                ciphered_data = downloaded_segments[-1][16:]

                cipher = AES.new(key, AES.MODE_CBC, iv=iv)
                data = unpad(cipher.decrypt(ciphered_data), AES.block_size)
                downloaded_segments[-1] = data

        return b''.join(downloaded_segments)

    @staticmethod
    def __download_key(key_uri: str) -> bin:
        return requests.get(url=key_uri).content
