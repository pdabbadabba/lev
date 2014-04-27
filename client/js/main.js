define(["event", "three"], function(el, three)
{
	
    var CONS = {
        // THREE.JS CONSTANTS
        // set the scene size
        WIDTH:904,
        HEIGHT:604,

        // set some camera attributes
        VIEW_ANGLE:45,
        NEAR:0.1,
        FAR:10000,

        CAMERA_X:0,
        CAMERA_Y:600,
        CAMERA_Z:1300
    }


    // setup default three.js stuff
    renderer = new THREE.WebGLRenderer();
    renderer.setSize(CONS.WIDTH, CONS.HEIGHT);
    renderer.setClearColor(0x0000ff);
    document.body.appendChild( renderer.domElement );

    camera = new THREE.PerspectiveCamera(CONS.VIEW_ANGLE, CONS.WIDTH / CONS.HEIGHT, CONS.NEAR, CONS.FAR);
    scene = new THREE.Scene();
    scene.add(camera);

    camera.position.z = CONS.CAMERA_Z;
    camera.position.x = CONS.CAMERA_X;
    camera.position.y = CONS.CAMERA_Y;
    camera.lookAt(scene.position);

    // add a light
    pointLight = new THREE.PointLight(0xFFFFFF);
    scene.add(pointLight);
    pointLight.position.x = 0;
    pointLight.position.y = 3000;
    pointLight.position.z = 0;
    pointLight.intensity = 8;

	

	var EventLoop = el
	eLoop = new EventLoop(scene, camera);

	(function mainLoop(){
		eLoop.tick((new Date()).getTime())
		requestAnimationFrame(mainLoop);
		renderer.render(scene, camera);
	})();
	
	//mainLoop.listen((new Date()).getTime(), 1000/30)

})