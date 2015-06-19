#!/usr/bin/env bash

# BEGIN subcommand functions

cmd_deps() {
    dep_type=${1:-ALL}
    # Treat all non-zero returns as fatal errors. Prevents issues like pip
    # package failing to install, causing the rest to be ignored.
    set -e
    case "$dep_type" in
      ALL)
        deps_pkgs ;;
      pkgs)
        deps_pkgs
        return ;;
    esac
    case "$dep_type" in
      ALL)
        deps_pip3 ;;
      pip)
        deps_pip3
        return ;;
    esac
    deps_misc
}

deps_pkgs() {
    if [ -e /etc/debian_version ]; then
        deps_debian
    else
        echo "As this is not a debian-based OS, please install the equivalents of these packages:"
        cat pkgs_debian.txt
    fi
}

deps_debian() {
    local pkgs=""
    echo "Checking for necessary debian packages"
    for pkg in $(< pkgs_debian.txt); do
        if ! dpkg-query -W --showformat='${Status}\n' $pkg 2> /dev/null | \
            grep -q "install ok installed"; then
            pkgs+="$pkg "
        fi
    done
    if [ -n "$pkgs" ]; then
        echo "Installing missing necessary packages: $pkgs"
        sudo DEBIAN_FRONTEND=noninteractive apt-get install $APTARGS --no-install-recommends $pkgs
    fi
}

deps_pip3() {
    echo "Installing necessary packages from pip3"
    pip3 install --user -r requirements.txt
}


deps_misc() {
    echo "Installing supervisor packages from pip2"
    pip2 install --user supervisor==3.1.3
    pip2 install --user supervisor-quick
}

cmd_init() {
    echo "Checking if tweetnacl has been built..."
    if [ -f lib/crypto/python-tweetnacl-20140309/build/python3.4/tweetnacl.so ] && [ -f lib/crypto/python-tweetnacl-20140309/build/python2.7/tweetnacl.so ]
    then
        echo "tweetnacl exists."
    else
        echo "tweetnacl.so does not exist. Compiling..."
        cd lib/crypto/python-tweetnacl-20140309/
        sh do
    fi
}

cmd_topology() {
    echo "Create topology, configuration, and execution files."
    mkdir -p logs traces
    PYTHONPATH=./ python3 topology/generator.py "$@"
}

cmd_setup() {
    echo "Add IP aliases for ISDs and ADs."
    for d in topology/ISD*; do
        for f in $d/setup/*; do
            sudo bash $f
        done
    done
}

cmd_run() {
    echo "Running the network..."
    supervisor/supervisor.sh reload
    supervisor/supervisor.sh quickstart all
}

cmd_stop() {
    echo "Terminating this run of the SCION infrastructure"
    supervisor/supervisor.sh quickstop all
}

cmd_clean() {
    {
    sudo ip addr flush dev lo
    sudo ip addr add 127.0.0.1/8 dev lo
    } &> /dev/null
    echo "Clean completed. Please check the output of ip addr to confirm the addresses were correctly flushed."
}

cmd_start(){
    # placeholder function to run all init functions
    # cmd_init
    # cmd_topology
    # cmd_setup
    # cmd_run
    echo "This method has not been fully implemented. Please run init, topology, setup, and run"
}

cmd_test(){
    PYTHONPATH=. nosetests -w test "$@"
}

cmd_coverage(){
    PYTHONPATH=. nosetests --with-cov -w test "$@"
    coverage html --omit 'external/*'
    echo "Coverage report here: file://$PWD/htmlcov/index.html"
}

cmd_version() {
	cat <<-_EOF
	============================================
	=                  SCION                   =
	=   https://github.com/netsec-ethz/scion   =
	============================================
	_EOF
}

cmd_help() {
	cmd_version
	echo
	cat <<-_EOF
	Usage:
	    $PROGRAM start
	        (not implemented) Performs all tasks (compile crypto lib, creates a topology, adds IP aliases, runs the network)
	    $PROGRAM deps
	        Install the necessary dependancies.
	    $PROGRAM init
	        Compile the SCION crypto library.
	    $PROGRAM topology
	        Create topology, configuration, and execution files.
	    $PROGRAM setup
	        Add IP aliases for ISDs and ADs.
	    $PROGRAM run
	        Run network.
	    $PROGRAM stop
	        Terminate this run of the SCION infrastructure.
	    $PROGRAM clean
	        Flush all the IP aliases of lo.
	    $PROGRAM test
	        Run all unit tests.
	    $PROGRAM coverage
	        Create a html report with unit test code coverage.
	    $PROGRAM help
	        Show this text.
	    $PROGRAM version
	        Show version information.
	_EOF
}
# END subcommand functions

PROGRAM="${0##*/}"
COMMAND="$1"
shift

case "$COMMAND" in
    clean|coverage|deps|help|init|run|setup|start|stop|test|topology|version)
        "cmd_$COMMAND" "$@" ;;
    *)  cmd_help ;;
esac
