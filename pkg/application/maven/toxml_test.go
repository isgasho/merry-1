package maven

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/merry/pkg/domain"
	"testing"
)

func Test_ShouldConvertToMavenXml(t *testing.T) {
	g := NewGomegaWithT(t)
	result := FromAntToXml(`
<project name="testing.example" default="dist">
    <modelVersion>4.0.0</modelVersion>
    <artifactId>my-test-app</artifactId>
    <groupId>my-test-group</groupId>
    <version>1.0-SNAPSHOT</version>

    <target name="test.management" depends="dist">
        <path id="classpatsh">
           <pathelement path="../../../_fixtures/jar/sample/annotations-13.0.jar" />
        </path>
    </target>
</project>
`, false, nil)

	g.Expect(result).To(Equal(`<?xml version="1.0"?>
<project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://maven.apache.org/POM/4.0.0"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

    <modelVersion>4.0.0</modelVersion>
	<groupId>my-test-group</groupId>        
	<artifactId>my-test-app</artifactId>
	<version>1.0-SNAPSHOT</version>
	<packaging>jar</packaging>

    <dependencies>
      <dependency>
          <groupId>org.jetbrains</groupId>
          <artifactId>annotations</artifactId>
          <version>13.0</version>
      </dependency>

    </dependencies>

</project>
`))
}


func Test_ShouldConvertToMavenXmlWithMap(t *testing.T) {
	g := NewGomegaWithT(t)

	depMap := make(map[string]domain.MavenDependency)
	depMap["org.jetbrains.annotations"] = domain.MavenDependency{
		Version:    "13.0.0",
		GroupId:    "org.jetbrains",
		ArtifactId: "jetbrains-annotations",
	}
	result := FromAntToXml(`
<project name="testing.example" default="dist">
    <modelVersion>4.0.0</modelVersion>
    <artifactId>my-test-app</artifactId>
    <groupId>my-test-group</groupId>
    <version>1.0-SNAPSHOT</version>

    <target name="test.management" depends="dist">
        <path id="classpatsh">
           <pathelement path="../../../_fixtures/jar/sample/annotations-13.0.jar" />
        </path>
    </target>
</project>
`, false, depMap)

	g.Expect(result).To(Equal(`<?xml version="1.0"?>
<project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://maven.apache.org/POM/4.0.0"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

    <modelVersion>4.0.0</modelVersion>
	<groupId>my-test-group</groupId>        
	<artifactId>my-test-app</artifactId>
	<version>1.0-SNAPSHOT</version>
	<packaging>jar</packaging>

    <dependencies>
      <dependency>
          <groupId>org.jetbrains</groupId>
          <artifactId>jetbrains-annotations</artifactId>
          <version>13.0.0</version>
      </dependency>

    </dependencies>

</project>
`))
}

