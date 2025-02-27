FROM alpine

ADD apache-traffic-server-exporter /usr/prometheus

WORKDIR /usr/prometheus

ENV ATS_LISTEN_PORT="default_exporter_addr"
ENV ATS_STATS_URI="default_ats_stats_uri"

CMD /usr/prometheus/apache-traffic-server-exporter
