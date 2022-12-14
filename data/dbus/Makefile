#
# Copyright (C) 2017 Canonical Ltd
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License version 3 as
# published by the Free Software Foundation.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

BINDIR := /usr/bin
LIBEXECDIR := /usr/lib
DBUSDIR = /usr/share/dbus-1
DBUSSERVICESDIR := ${DBUSDIR}/services

SERVICES_GENERATED := $(patsubst %.service.in,%.service,$(wildcard *.service.in))
SERVICES := ${SERVICES_GENERATED}

%.service: %.service.in
	cat $< | \
		sed s:@libexecdir@:$(LIBEXECDIR):g | \
		sed 's:@bindir@:${BINDIR}:g' | \
		cat > $@

all: ${SERVICES}

install:: ${SERVICES}
	# NOTE: old (e.g. 14.04) GNU coreutils doesn't -D with -t
	install -d -m 0755 ${DESTDIR}/${DBUSSERVICESDIR}
	install -m 0644 -t ${DESTDIR}/${DBUSSERVICESDIR} $^

install:: snapd.session-services.conf
	install -d -m 0755 ${DESTDIR}/${DBUSDIR}/session.d
	install -m 0644 -t ${DESTDIR}/${DBUSDIR}/session.d $^

install:: snapd.system-services.conf
	install -d -m 0755 ${DESTDIR}/${DBUSDIR}/system.d
	install -m 0644 -t ${DESTDIR}/${DBUSDIR}/system.d $^

clean:
	rm -f ${SERVICES_GENERATED}
