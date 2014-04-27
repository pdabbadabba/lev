define(
	function() {

	function Terrain(x, y) {

		this.x = x;
		this.y = y;

		this.addToScene = function(scene)
		{
			n = 0
		    function loaded() {
		        n++;
		        console.log("loaded: " + n);
		 
		        if (n == 3) {
		            terrain.visible = true;
		            console.log('ff');
		            scene.add(terrain);
		        }
		    }

		      // load the heightmap we created as a texture
	          var texture = THREE.ImageUtils.loadTexture('http://localhost:8080/png?x='+x+'&y='+y, null, loaded);

              // texture two other textures we'll use to make the map look more real
              var detailTexture = THREE.ImageUtils.loadTexture("http://localhost:8000/assets/bg.jpg", null, loaded);


              // the following configuration defines how the terrain is rendered
              var terrainShader = THREE.ShaderTerrain[ "terrain" ];
              var uniformsTerrain = THREE.UniformsUtils.clone(terrainShader.uniforms);

              // how to treat abd scale the normal texture
              uniformsTerrain[ "tNormal" ].texture = detailTexture;
              uniformsTerrain[ "uNormalScale" ].value = 1;

	          // the displacement determines the height of a vector, mapped to
              // the heightmap
              uniformsTerrain[ "tDisplacement" ].texture = texture;
              uniformsTerrain[ "uDisplacementScale" ].value = 100;

              // the following textures can be use to finetune how
              // the map is shown. These are good defaults for simple
              // rendering
              uniformsTerrain[ "tDiffuse1" ].texture = detailTexture;
              uniformsTerrain[ "tDetail" ].texture = detailTexture;
              uniformsTerrain[ "enableDiffuse1" ].value = true;
              uniformsTerrain[ "enableDiffuse2" ].value = true;
              uniformsTerrain[ "enableSpecular" ].value = true;

              // // diffuse is based on the light reflection
              uniformsTerrain[ "uDiffuseColor" ].value.setHex(0xcccccc);
              uniformsTerrain[ "uSpecularColor" ].value.setHex(0xff0000);
              // is the base color of the terrain
              uniformsTerrain[ "uAmbientColor" ].value.setHex(0x0000cc);

              // // how shiny is the terrain
              uniformsTerrain[ "uShininess" ].value = 3;

              // // handles light reflection
              uniformsTerrain[ "uRepeatOverlay" ].value.set(6, 6);

              // configure the material that reflects our terrain
              var material = new THREE.ShaderMaterial({
                  uniforms:uniformsTerrain,
                  vertexShader:terrainShader.vertexShader,
                  fragmentShader:terrainShader.fragmentShader,
                  lights:true,
                  fog:false
              });

              // we use a plain to render as terrain
              var geometryTerrain = new THREE.PlaneGeometry(2000, 2000, 256, 256);
              geometryTerrain.applyMatrix(new THREE.Matrix4().makeRotationX(Math.PI / 2));
              geometryTerrain.computeFaceNormals();
              geometryTerrain.computeVertexNormals();
              geometryTerrain.computeTangents();

              // create a 3D object to add
              terrain = new THREE.Mesh(geometryTerrain, material);
              terrain.position.set(0, -125, -4000);
              terrain.rotation.x = -Math.PI / 2;
              terrain.visible = true;

              loaded()
		}


	}

	return Terrain;
})