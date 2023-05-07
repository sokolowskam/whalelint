import org.jetbrains.kotlin.gradle.tasks.KotlinCompile
import org.jetbrains.kotlin.platform.TargetPlatformVersion.NoVersion.description

plugins {
    id("org.jetbrains.intellij") version "1.13.3"
    id("org.jetbrains.kotlin.jvm") version "1.8.21"
    id("org.kordamp.gradle.markdown") version "2.2.0"
    id("java")
}

group = "WhaleLint"
version = "0.0.7"

description = "WhaleLint is a Dockerfile linter written in Golang."

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib-jdk8"))
    implementation("com.google.code.gson:gson:2.8.6" )
    implementation("org.jetbrains:annotations:20.1.0")
    runtimeOnly(group = "commons-io", name = "commons-io", version = "2.6")
}

// See https://github.com/JetBrains/gradle-intellij-plugin/
intellij {
    version.set("2023.1")
    pluginName.set("whalelint")
    updateSinceUntilBuild.set(false)
    // setPlugins("Docker:$version")
}

tasks.register<Copy>("copyChangelogAndReadme") {
    from(file("$buildDir/../../vscode/CHANGELOG.md"), file("readme.md"))
    into(file("$buildDir/idea-sandbox/plugins/whalelint/docs/"))
}

tasks.markdownToHtml {
    sourceDir = file("$buildDir/idea-sandbox/plugins/whalelint/docs")
    outputDir = file("$buildDir/idea-sandbox/plugins/whalelint/docs")
    dependsOn("copyChangelogAndReadme")
}

tasks.getByName<org.jetbrains.intellij.tasks.PatchPluginXmlTask>("patchPluginXml") {
    dependsOn("copyChangelogAndReadme", "markdownToHtml")

    val changelogPath = "$buildDir/idea-sandbox/plugins/whalelint/docs/CHANGELOG.html"
    val readmePath    = "$buildDir/idea-sandbox/plugins/whalelint/docs/readme.html"


    if (file(changelogPath).exists()) {
        changeNotes.set(file(changelogPath).readText())
        //changeNotes(file(changelogPath).readText())
    }
    if (file(readmePath).exists()) {
        pluginDescription.set(file(readmePath).readText().replace(
                "<h1>WhaleLint JetBrains Plugin</h1>", "").replace(
                "<h2>Introduction</h2>", ""))
    }

    version.set("0.0.7")
}

tasks.withType<JavaCompile> {
    sourceCompatibility = "1.8"
    targetCompatibility = "1.8"
}

listOf("compileKotlin", "compileTestKotlin").forEach {
    tasks.getByName<KotlinCompile>(it) {
        kotlinOptions.jvmTarget = "1.8"
    }
}

tasks.register("copyWhaleLintBinary") {
    doLast {
        copy {
            from("$projectDir/../../../whalelint/whalelint")
            into("$buildDir/idea-sandbox/plugins/whalelint/bin/")
        }
    }
}

tasks.named("prepareSandbox") {
    finalizedBy("copyWhaleLintBinary")


}

tasks.buildPlugin {
    dependsOn("copyChangelogAndReadme", "markdownToHtml")
}

tasks.publishPlugin {
    dependsOn("copyChangelogAndReadme", "markdownToHtml")
    token.set(System.getenv("JETBRAINS_TOKEN"))
}
