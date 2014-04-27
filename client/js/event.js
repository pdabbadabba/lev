define(["terrain"],
	function(terrain) {

	function EventLoop(scene, camera){


	this.scene = scene
	this.camera = camera



	this.terr = new terrain(0,0)

	var geometry = new THREE.CubeGeometry(500, 500, 500);
	var material = new THREE.MeshLambertMaterial( { color: 0x00ff00 } );
	this.cube = new THREE.Mesh( geometry, material );
	this.cube.position.set(0, -125, -4000)
	

	this.scene.add( this.cube );

	this.terr.addToScene(scene)

	this.camera.position.z = 5;
	this.prev_ticks = (new Date()).getTime()

		this.tick = function(ticks) {

			msDelta = ticks - this.prev_ticks;
			
        	// camera.position.x = (Math.cos( msDelta * 0.01 ) *  1300);
        	// camera.position.z = (Math.sin( msDelta * 0.01) *  1300) ;

			this.cube.rotation.x += 0.001 * msDelta;
			this.cube.rotation.y += 0.001 * msDelta;

			this.prev_ticks = ticks
		}		
	}

	return EventLoop;

	}



	
)