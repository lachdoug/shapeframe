# options: 
#  -c to continue after failed test
#  -d to execute command in debug mode

export PATH=$PATH:.

redColor='\033[0;31m'
greenColor='\033[0;32m'
yellowColor='\033[0;33m'
blueColor='\033[0;34m'
resetText='\033[0m'

errorCount=0
successCount=0
args=( "$@" )

sf() {
    echo -e "${blueColor}sf $@${resetText}"
    if [[ $args =~ "-d" ]]
    then
        command sf --debug "$@"
    else
        command sf "$@"
    fi    

    if [ $? != 0 ]
    then
        handleFailure
    else
        handleSuccess
    fi
}

handleFailure() {
    echo -e "${redColor}Fail${resetText}\n"
    let errorCount=errorCount+1
    if ! [[ $args =~ "-c" ]]
    then
        exit 1
    fi    
}

handleSuccess() {
    echo -e "${greenColor}Pass${resetText}\n"
    let successCount=successCount+1
}

report() {
    echo -e "${greenColor}Pass: ${successCount} ${redColor}Fail: ${errorCount}${resetText}"
    if [[ "$errorCount" -eq 0 ]]
    then
        echo -e "${greenColor}Tests passed${resetText}"
        exit 0
    else
        echo -e "${redColor}Tests failed${resetText}"
        exit 1
    fi
}

t() {
    echo -e "\n${yellowColor}$@${resetText}\n"   
}

nuke() {
    echo "nuke workspace"
    rm sf.log
    rm sf.db
    rm sf.db.log
    rm -rf repos
}

# Config settings

frame_config_yaml="
Host: 0.0.0.0
"
shape_config_yaml="
color: red
"
frameshape_config_yaml="
restart: on-failure
restart-on-failure: 3
"

t Clean slate
nuke

t Initialize
sf init
sf g w
nuke
sf init -name Test -about "Workspace test"
sf g w
nuke

t Adding and removing directories
sf init
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf rm d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
nuke

t Adding and removing repositories
sf init
sf a r github.com/lachdoug/shapeframe-apps
sf pull github.com/lachdoug/shapeframe-apps
sf rm r github.com/lachdoug/shapeframe-apps
sf a r -https github.com/lachdoug/shapeframe-apps
sf pull github.com/lachdoug/shapeframe-apps
sf checkout github.com/lachdoug/shapeframe-apps main
nuke

t Adding and removing frames
sf init
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf rm f docker-local
sf a f docker-local
sf cx docker-local
sf rm f
sf cx
nuke

t Adding and removing shapes
sf init
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf cx docker-local
sf a s apache
sf rm s apache
sf cx ..
sf a s -frame docker-local apache
sf rm s -frame docker-local apache
nuke

t Listing
sf init
sf a r github.com/lachdoug/shapeframe-apps
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf a s -frame docker-local apache
sf ls f
sf ls s
sf ls sr
sf ls fr
nuke

t Reading
sf init
sf a r github.com/lachdoug/shapeframe-apps
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf a s -frame docker-local apache
sf g w
sf g f docker-local
sf g s -frame docker-local apache
sf cx docker-local
sf g f
sf g s apache
sf cx apache
sf g s
nuke

t Inspecting
sf init
sf a r github.com/lachdoug/shapeframe-apps
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf a s -frame docker-local apache
sf i
nuke

t Contexting
sf init
sf cx
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf a s -frame docker-local apache
sf cx docker-local
sf cx
sf cx apache
sf cx
nuke

t Configuring
sf init
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf a s -frame docker-local apache
sf cg f -frame docker-local "$frame_config_yaml"
sf cg s -frame docker-local -shape apache "$shape_config_yaml"
sf cg s-f -frame docker-local -shape apache "$frameshape_config_yaml"
sf cx docker-local
sf cg s -shape apache "$shape_config_yaml"
sf cg s-f -shape apache "$frameshape_config_yaml"
sf cx apache
sf cg s "$shape_config_yaml"
sf cg s-f "$frameshape_config_yaml"
nuke

t Labelling
sf init
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf a s -frame docker-local apache
sf ll f -name docker-localA -about "Deploy containers on local Docker EngineA" docker-local
sf ll s -frame docker-localA -name apacheA -about "Apache (Web Server)A" apache
sf ll f -name docker-localB -about "Deploy containers on local Docker EngineB" docker-localA
sf ll s -frame docker-localB -name apacheB -about "Apache (Web Server)B" apacheA
sf cx docker-localB
sf ll f -name docker-localC -about "Deploy containers on local Docker EngineC"
sf ll s -name apacheC -about "Apache (Web Server)C" apacheB
sf cx apacheC
sf ll s -name apacheD -about "Apache (Web Server)D"
sf g f
sf g s 
nuke

t Orchestrating
sf init
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf a s -frame docker-local apache
sf cg f -frame docker-local "$frame_config_yaml"
sf cg s -frame docker-local -shape apache "$shape_config_yaml"
sf cg s-f -frame docker-local -shape apache "$frameshape_config_yaml"
sf o -frame docker-local
sf cx docker-local
sf o

t Helping
sf -help
sf initialize -help
sf list -help
sf ls frames -help
sf ls shapes -help
sf ls framers -help
sf ls shapers -help
sf label -help
sf ll shape -help
sf ll frame -help
sf ll workspace -help
sf configure -help
sf cg shape -help
sf cg frame -help
sf inspect -help
sf add -help
sf a shape -help
sf a frame -help
sf a repository -help
sf a directory -help
sf remove -help
sf rm shape -help
sf rm frame -help
sf rm repository -help
sf rm directory -help
sf pull -help
sf context -help
sf orchestrate -help

report
